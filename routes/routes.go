package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/controllers"
	"backend/middleware"
	"log"
)

func SetupRoutes(app *fiber.App) {
	// Root Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Full-Stack Go & Angular Authentication API!")
	})

	// Example private route
	app.Get("/api/private/data", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		username := c.Locals("username")
		if username == nil {
			log.Println("Username not found in context")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
		}
		log.Println("Accessing private data for user:", username.(string))
		return c.JSON(fiber.Map{"message": "Welcome, " + username.(string)})
	})
	
	// API Routes
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/refresh", controllers.RefreshToken)


	// Protected routes (using the updated AuthMiddleware)
	app.Use("/api/private", middleware.AuthMiddleware)
	app.Get("/api/protected", middleware.AuthMiddleware, controllers.ProtectedHandler)

}
