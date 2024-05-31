package main

import (
    "log"
    "os"

    "github.com/AquibPy/shorten-url-fiber-redis/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"
    _ "github.com/AquibPy/shorten-url-fiber-redis/docs" // import generated docs
    "github.com/gofiber/swagger"                        // swagger handler
)

// @title Fiber URL Shortener API
// @version 1.0
// @description This is a sample URL shortener server.
// @host localhost:3000
// @BasePath /
func setupRoutes(app *fiber.App) {
    app.Get("/:url", routes.ResolveURL)
    app.Post("/api/v1", routes.ShortenURL)
    app.Get("/swagger/*", swagger.HandlerDefault) // default
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    app := fiber.New()

    // Add CORS middleware
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*", // Allow all origins
        AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

    app.Use(logger.New())
    setupRoutes(app)
    log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
