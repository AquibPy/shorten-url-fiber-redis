# URL Shortener with Go Fiber and Redis

In a digital landscape inundated with long and cumbersome URLs, a URL shortener service comes as a handy tool to condense links into more manageable forms. This project presents a URL shortener service built with the Go Fiber framework and Redis, designed to simplify URL management and enhance user experience.

### Motivation

The motivation behind this project stems from the need for a lightweight, efficient, and scalable solution to shorten URLs without compromising on features or security. Whether for sharing links on social media, optimizing space in printed materials, or tracking clicks for analytics purposes, a URL shortener offers a versatile solution for various use cases.

### Potential Use Cases

- **Social Media Sharing**: Simplify sharing of links on platforms like Twitter where character limits are stringent.
- **Email Campaigns**: Create concise and visually appealing links for inclusion in email newsletters or marketing campaigns.
- **Printed Materials**: Reduce lengthy URLs to QR codes or short, memorable links for print advertisements, flyers, or business cards.
- **Affiliate Marketing**: Mask affiliate URLs to make them more aesthetically pleasing and increase click-through rates.
- **Analytics Tracking**: Monitor click-through rates, geographical distribution of clicks, and other analytics data for marketing purposes.
- **Internal Link Management**: Streamline internal processes by shortening and organizing links to company resources, documents, or tools.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Swagger Documentation](#swagger-documentation)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Docker](#docker)
- [Contributing](#contributing)

## Introduction

This is a URL shortener service built with the Go Fiber framework and Redis. It provides a simple API to shorten URLs and resolve shortened URLs. The service includes rate limiting and URL validation features.

## Features

- Shorten URLs with optional custom aliases
- Resolve shortened URLs to their original form
- Enforce HTTP scheme on URLs
- Basic rate limiting to prevent abuse
- Dockerized for easy deployment
- Swagger documentation for easy API exploration

## Prerequisites

- [Go](https://golang.org/dl/) (1.16 or higher)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/AquibPy/shorten-url-fiber-redis.git
    cd shorten-url-fiber-redis
    ```

2. Create a `.env` file in the root directory with the necessary environment variables (refer to the [Environment Variables](#environment-variables) section).

3. Build and run the Docker containers:
    ```bash
    docker-compose up --build
    ```

## Usage

After running the Docker containers, the API will be available at `http://localhost:3000`.

## API Endpoints

### Shorten URL

- **Endpoint**: `POST /api/v1`
- **Description**: Shorten a URL with an optional custom alias.
- **Request Body**:
    ```json
    {
        "url": "https://example.com",
        "short": "customAlias",
        "expiry": 24
    }
    ```
- **Response**:
    ```json
    {
        "url": "https://example.com",
        "short": "http://localhost:3000/customAlias",
        "expiry": 24,
        "rate_limit": 10,
        "rate_limit_reset": 30
    }
    ```

### Resolve URL

- **Endpoint**: `GET /:url`
- **Description**: Resolve a shortened URL to its original form.
- **Response**:
    - Redirects to the original URL.
    - Returns 404 if the shortened URL is not found.

## Swagger Documentation

Swagger documentation is available to explore the API interactively.

### Setup

1. Install the `swag` tool for generating Swagger docs:
    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

2. Generate the Swagger documentation:
    ```bash
    swag init
    ```

3. Access the Swagger UI at `http://localhost:3000/swagger/index.html`.

### Example

Here is how to integrate Swagger in your `main.go`:

```go
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
```

## Project Structure

```
.
├── api
│   ├── Dockerfile
│   ├── database
│   │   └── database.go
│   ├── helpers
│   │   └── helpers.go
│   ├── routes
│   │   ├── resolve.go
│   │   └── shorten.go
│   ├── main.go
│   ├── .env
│   └── docs
│       └── swagger.json
├── db
│   └── Dockerfile
├── docker-compose.yml
└── README.md
```

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```
APP_PORT=3000
DB_ADDR=db:6379
DB_PASS=
DOMAIN=http://localhost:3000
API_QUOTA=10
```

## Docker

### API Dockerfile

The API Dockerfile uses a multistage build to reduce the final image size:

1. Build the Go application in a `golang:alpine` container.
2. Copy the built binary to a smaller `alpine` container for deployment.

### Database Dockerfile

The database Dockerfile uses the official Redis image.

### Docker Compose

The `docker-compose.yml` file defines the services:

- `api`: The URL shortener API service.
- `db`: The Redis database service.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your changes.