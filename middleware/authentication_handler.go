package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt"
    "backend/utils"
    "log"
)

func AuthMiddleware(c *fiber.Ctx) error {
    // Get the token from cookies
    tokenString := c.Cookies("access_token")
    if tokenString == "" {
        log.Println("Access token missing from cookies")
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Access token missing"})
    }
    log.Println("Token from cookie:", tokenString)

    // Parse and validate the JWT token
    token, err := jwt.ParseWithClaims(tokenString, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })

    if err != nil || !token.Valid {
        log.Println("Token parsing error:", err)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
    }

    // Extract claims
    claims, ok := token.Claims.(*utils.CustomClaims)
    if !ok || claims.Username == "" {
        log.Println("Invalid token claims or username empty")
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
    }

    // Log the extracted username for verification
    log.Println("Authenticated user from token claims:", claims.Username)

    // Set the username in the context
    c.Locals("username", claims.Username)
    return c.Next()
}
