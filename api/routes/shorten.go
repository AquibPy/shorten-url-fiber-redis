package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/AquibPy/shorten-url-fiber-redis/database"
	"github.com/AquibPy/shorten-url-fiber-redis/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// ShortenURL godoc
// @Summary Shortens a URL
// @Description Creates a shortened version of a given URL.
// @Tags URL
// @Accept json
// @Produce json
// @Param request body ShortenURLRequest true "Request body"
// @Success 200 {object} ShortenURLResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1 [post]
func ShortenURL(c *fiber.Ctx) error {
	body := new(ShortenURLRequest)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(ErrorResponse{Error: "cannot parse json"})
	}

	r2 := database.CreateClient(1)
	defer r2.Close()
	val, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(RateLimitExceededResponse{
				Error:          "Rate limit exceeded",
				RateLimitReset: Expiry(limit / time.Nanosecond / time.Minute),
			})
		}
	}

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid URL"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(ErrorResponse{
			Error: "haha... nice try",
		})
	}

	body.URL = helpers.EnforceHTTP(body.URL)

	var id string

	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)

	val, _ = r.Get(database.Ctx, id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(ErrorResponse{
			Error: "URL custom short is already in use",
		})
	}

	if body.Expiry == 0 {
		body.Expiry = 24 // default expiry of 24 hours
	}

	err = r.Set(database.Ctx, id, body.URL, time.Duration(body.Expiry)*3600*time.Second).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Unable to connect to server",
		})
	}

	resp := ShortenURLResponse{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemainig:   10,
		XRateLimitReset: 30,
	}

	r2.Decr(database.Ctx, c.IP())
	val, _ = r2.Get(database.Ctx, c.IP()).Result()

	resp.XRateRemainig, _ = strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = Expiry(ttl / time.Nanosecond / time.Minute)

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id
	return c.Status(fiber.StatusOK).JSON(resp)
}
