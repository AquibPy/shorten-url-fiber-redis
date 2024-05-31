package routes

import (
    "github.com/AquibPy/shorten-url-fiber-redis/database"
    "github.com/gofiber/fiber/v2"
    "github.com/redis/go-redis/v9"
)

// ResolveURL godoc
// @Summary Resolves a shortened URL
// @Description Redirects to the original URL corresponding to the given short URL.
// @Tags URL
// @Produce json
// @Param url path string true "Short URL"
// @Success 301
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /{url} [get]
func ResolveURL(c *fiber.Ctx) error {
    url := c.Params("url")

    r := database.CreateClient(0)
    defer r.Close()

    value, err := r.Get(database.Ctx, url).Result()

    if err == redis.Nil {
        return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
            Error: "short not found on database",
        })
    } else if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
            Error: "cannot connect to DB",
        })
    }

    rInr := database.CreateClient(1)
    defer rInr.Close()

    _ = rInr.Incr(database.Ctx, "counter")
    return c.Redirect(value, 301)
}
