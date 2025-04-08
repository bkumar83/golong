package tests
import (
	"testing"
	"backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

// TestCORSSetup tests the CORS middleware for proper headers
func TestCORSSetup(t *testing.T) {
	app := fiber.New()
	middleware.SetupCORS(app)

	app.Get("/api/test", func(c *fiber.Ctx) error {
		return c.SendString("CORS Check")
	})

	req := httptest.NewRequest("GET", "/api/test", nil)
	req.Header.Set("Origin", "http://localhost:4200")
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "http://localhost:4200", resp.Header.Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", resp.Header.Get("Access-Control-Allow-Credentials"))
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Methods"), "GET")
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Methods"), "POST")
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Methods"), "PUT")
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Methods"), "DELETE")
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Headers"), "Content-Type")
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Headers"), "Authorization")
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Headers"), "Accept")
	assert.Equal(t, "Set-Cookie", resp.Header.Get("Access-Control-Expose-Headers"))
}