package routes

// ErrorResponse represents a common structure for error responses
type ErrorResponse struct {
    Error string `json:"error"`
}

// RateLimitExceededResponse represents the response when the rate limit is exceeded
type RateLimitExceededResponse struct {
    Error           string `json:"error"`
    RateLimitReset  Expiry `json:"rate_limit_reset"`
}

// Expiry represents a duration in hours as an integer
type Expiry int

// ShortenURLRequest represents the payload for shortening a URL
type ShortenURLRequest struct {
    URL         string `json:"url"`
    CustomShort string `json:"short"`
    Expiry      Expiry `json:"expiry"`
}

// ShortenURLResponse represents the response after a URL is shortened
type ShortenURLResponse struct {
    URL             string `json:"url"`
    CustomShort     string `json:"short"`
    Expiry          Expiry `json:"expiry"`
    XRateRemainig   int    `json:"rate_limit"`
    XRateLimitReset Expiry `json:"rate_limit_reset"`
}
