// Authentication controller
// Auth Controller (auth_controller.go)
package controllers

import (
    "github.com/gofiber/fiber/v2"
    "backend/config"
    "backend/models"
    "backend/utils"
    "log"
)

// Register Handler
// Register Handler
func Register(c *fiber.Ctx) error {
    user := new(models.User)

    // Parse the request body
    if err := c.BodyParser(user); err != nil {
        log.Println("Error parsing registration request body:", err)
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    log.Println("Received Registration Username:", user.Username)
    log.Println("Received Registration Password:", user.Password)

    // Hash the password before storing
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        log.Println("Error hashing password during registration:", err)
        return c.Status(500).JSON(fiber.Map{"error": "Password hashing failed"})
    }

    user.Password = hashedPassword

    // Save the user to the database
    if err := config.DB.Create(&user).Error; err != nil {
        log.Println("Error during registration:", err)
        return c.Status(500).JSON(fiber.Map{"error": "Registration failed"})
    }

    log.Println("User registered successfully:", user.Username)
    return c.JSON(fiber.Map{"message": "User registered successfully"})
}




// Login Handler
// Login Handler
func Login(c *fiber.Ctx) error {
    var user models.User
    var dbUser models.User

    if err := c.BodyParser(&user); err != nil {
        log.Println("Error parsing login request body:", err)
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    if config.IsKratosEnabled() {
        // Ory Kratos login flow
        isValid, err := utils.ValidateKratosSession(c.Cookies("access_token"))
        if err != nil || !isValid {
            log.Println("Kratos authentication failed:", err)
            return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
        }
        log.Println("Authenticated with Ory Kratos")
    } else {
        // Legacy authentication
        if err := config.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
            return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
        }
        if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
            return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
        }
    }

    // Generate JWT tokens
    accessToken, err := utils.GenerateJWT(user.Username)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to generate access token"})
    }

    refreshToken, err := utils.GenerateRefreshToken(user.Username)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to generate refresh token"})
    }

    // Set refresh token as HTTP-only cookie
    c.Cookie(&fiber.Cookie{
        Name:     "refresh_token",
        Value:    refreshToken,
        HTTPOnly: true,
        Secure:   true,
        Path:     "/",
    })

    return c.JSON(fiber.Map{
        "access_token":  accessToken,
        "refresh_token": refreshToken,
    })
}


func ProtectedHandler(c *fiber.Ctx) error {
    username := c.Locals("username")
    if username == nil {
        log.Println("Username not found in context")
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
    }
    log.Println("Accessing protected data for user:", username.(string))
    return c.JSON(fiber.Map{
        "message": "Welcome to protected data, " + username.(string),
    })
}

// Refresh Token Handler
func RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Refresh token missing"})
	}

	// Validate the refresh token using middleware
	username, err := utils.ExtractUsernameFromToken(refreshToken)
	if err != nil {
		log.Println("Invalid refresh token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	// Generate a new access token
	newAccessToken, err := utils.GenerateJWT(username)
	if err != nil {
		log.Println("Error generating new access token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token generation failed"})
	}

	return c.JSON(fiber.Map{"access_token": newAccessToken})
}
