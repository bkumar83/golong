// Rate limiting middleware
package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2"
	"time"
)

func SetupRateLimiter(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
	}))
}
