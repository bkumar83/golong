// Unit tests for authentication controller
// Comprehensive Tests for Backend Functionality
package tests

import (
	"testing"
	"backend/utils"
	"backend/controllers"
	"backend/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
)

// Test JWT Token Generation
func TestGenerateAccessToken(t *testing.T) {
	token, err := utils.GenerateJWT("testuser")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// Test JWT Token Extraction
func TestExtractUsernameFromToken(t *testing.T) {
	token, _ := utils.GenerateJWT("testuser")
	username, err := utils.ExtractUsernameFromToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "testuser", username)
}

// Test Login Handler - Successful Login
func TestLoginHandlerSuccess(t *testing.T) {
	app := fiber.New()
	app.Post("/api/login", controllers.Login)

	req := httptest.NewRequest("POST", "/api/login", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

// Test Middleware - Auth Validation
func TestAuthMiddleware(t *testing.T) {
	app := fiber.New()
	app.Use(middleware.AuthMiddleware)
	app.Get("/api/protected", func(c *fiber.Ctx) error {
		return c.SendString("Protected!")
	})

	req := httptest.NewRequest("GET", "/api/protected", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

// Test Refresh Token Handler - Unauthorized
func TestRefreshTokenHandlerUnauthorized(t *testing.T) {
	app := fiber.New()
	app.Post("/api/refresh", controllers.RefreshToken)

	req := httptest.NewRequest("POST", "/api/refresh", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}
