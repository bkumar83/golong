// CORS middleware
package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

func SetupCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:4200",  // Specify your frontend origin
        AllowCredentials: true,                 // Allow credentials
        AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        ExposeHeaders:    "Set-Cookie",              // Expose the cookie header
    }))
}

