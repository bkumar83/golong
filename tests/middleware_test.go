// Unit tests for middleware
package tests

import (
	"testing"
	"backend/middleware"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/http"
)

// TestAuthMiddleware_ValidToken tests middleware with a valid token
func TestAuthMiddleware_ValidToken(t *testing.T) {
	app := fiber.New()
	app.Use(middleware.AuthMiddleware)
	app.Get("/api/protected", func(c *fiber.Ctx) error {
		return c.SendString("Protected!")
	})

	// Generate a valid token
	token, _ := utils.GenerateJWT("testuser")

	req := httptest.NewRequest("GET", "/api/protected", nil)
	req.AddCookie(&http.Cookie{Name: "access_token", Value: token})

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

// TestAuthMiddleware_MissingToken tests middleware with a missing token
func TestAuthMiddleware_MissingToken(t *testing.T) {
	app := fiber.New()
	app.Use(middleware.AuthMiddleware)
	app.Get("/api/protected", func(c *fiber.Ctx) error {
		return c.SendString("Protected!")
	})

	req := httptest.NewRequest("GET", "/api/protected", nil)

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

// TestAuthMiddleware_InvalidToken tests middleware with an invalid token
func TestAuthMiddleware_InvalidToken(t *testing.T) {
	app := fiber.New()
	app.Use(middleware.AuthMiddleware)
	app.Get("/api/protected", func(c *fiber.Ctx) error {
		return c.SendString("Protected!")
	})

	req := httptest.NewRequest("GET", "/api/protected", nil)
	req.AddCookie(&http.Cookie{Name: "access_token", Value: "invalidtoken"})

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}
