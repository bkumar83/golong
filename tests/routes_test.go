// Unit tests for routes setup
package tests

import (
	"testing"
	"backend/routes"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

// TestRootRoute tests the root route for a welcome message
func TestRootRoute(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestPrivateDataRoute_ValidToken tests private data access with a valid token
func TestPrivateDataRoute_ValidToken(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	token, _ := utils.GenerateJWT("testuser")
	req := httptest.NewRequest("GET", "/api/private/data", nil)
	req.AddCookie(&http.Cookie{Name: "access_token", Value: token})

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestPrivateDataRoute_MissingToken tests private data access without a token
func TestPrivateDataRoute_MissingToken(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	req := httptest.NewRequest("GET", "/api/private/data", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

// TestProtectedRoute_ValidToken tests protected route access with a valid token
func TestProtectedRoute_ValidToken(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	token, _ := utils.GenerateJWT("testuser")
	req := httptest.NewRequest("GET", "/api/protected", nil)
	req.AddCookie(&http.Cookie{Name: "access_token", Value: token})

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestProtectedRoute_MissingToken tests protected route access without a token
func TestProtectedRoute_MissingToken(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	req := httptest.NewRequest("GET", "/api/protected", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

// TestLoginRoute_InvalidCredentials tests the login route with invalid credentials
func TestLoginRoute_InvalidCredentials(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	req := httptest.NewRequest("POST", "/api/login", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
