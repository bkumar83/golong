package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"backend/config"
	"backend/routes"
	"backend/middleware"
	"backend/utils"     
	"io"
	"log"
	"os"
)

func main() {
    // Load environment variables
    config.LoadEnv()

    // Initialize Ory Kratos client
    utils.InitKratosClient(os.Getenv("ORY_KRATOS_URL"))

    // Connect to the database
    config.ConnectDB()

    // Create a new log file or open it if it exists
    file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal("Error opening log file:", err)
    }
    defer file.Close()

    // Log to both file and console using MultiWriter
    multiWriter := io.MultiWriter(os.Stdout, file)
    log.SetOutput(multiWriter)

    // Create a new Fiber instance
    app := fiber.New()

    // Enable request logging using Fiber's logger middleware
    app.Use(logger.New(logger.Config{
        Format:     "[${time}] ${status} - ${method} ${path} - ${latency}\n",
        TimeFormat: "02-Jan-2006 15:04:05",
        TimeZone:   "Local",
        Output:     multiWriter,
    }))

    // Log server startup
    log.Println("Server starting on http://localhost:8080")

    // Setup middleware
    middleware.SetupCORS(app)
    middleware.SetupRateLimiter(app)

    // Setup API routes
    routes.SetupRoutes(app)

    // Start the server
    if err := app.Listen(":8080"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}

