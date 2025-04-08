// Unit tests for JWT utility functions
package tests

import (
	"testing"
	"backend/utils"
	"github.com/stretchr/testify/assert"
	"log"
)

// TestGenerateJWT_Success tests successful JWT generation
func TestGenerateJWT_Success(t *testing.T) {
	username := "testuser"
	token, err := utils.GenerateJWT(username)

	// Assertions
	assert.NoError(t, err, "Error should be nil when generating JWT")
	assert.NotEmpty(t, token, "Generated token should not be empty")
	log.Println("Generated JWT:", token)
}

// TestGenerateJWT_EmptyUsername tests generating a token with an empty username
func TestGenerateJWT_EmptyUsername(t *testing.T) {
	username := ""
	token, err := utils.GenerateJWT(username)

	// Assertions
	assert.NoError(t, err, "Error should be nil even with an empty username")
	assert.NotEmpty(t, token, "Generated token should not be empty")
	log.Println("Generated JWT with empty username:", token)
}

// TestGenerateRefreshToken_Success tests successful refresh token generation
func TestGenerateRefreshToken_Success(t *testing.T) {
	username := "testuser"
	refreshToken, err := utils.GenerateRefreshToken(username)

	// Assertions
	assert.NoError(t, err, "Error should be nil when generating refresh token")
	assert.NotEmpty(t, refreshToken, "Generated refresh token should not be empty")
	log.Println("Generated Refresh Token:", refreshToken)
}

// TestExtractUsernameFromToken_Success tests extracting the username from a valid token
func TestExtractUsernameFromToken_Success(t *testing.T) {
	username := "testuser"
	token, _ := utils.GenerateJWT(username)

	extractedUsername, err := utils.ExtractUsernameFromToken(token)

	// Assertions
	assert.NoError(t, err, "Error should be nil when extracting username from a valid token")
	assert.Equal(t, username, extractedUsername, "Extracted username should match the original")
	log.Println("Extracted Username:", extractedUsername)
}

// TestExtractUsernameFromToken_Invalid tests extracting from an invalid token
func TestExtractUsernameFromToken_Invalid(t *testing.T) {
	invalidToken := "invalidtoken"
	username, err := utils.ExtractUsernameFromToken(invalidToken)

	// Assertions
	assert.Error(t, err, "Error should occur when extracting from an invalid token")
	assert.Equal(t, "", username, "Username should be empty for an invalid token")
	log.Println("Failed to extract username from invalid token as expected")
}

// TestExtractUsernameFromToken_Empty tests extracting from an empty token
func TestExtractUsernameFromToken_Empty(t *testing.T) {
	emptyToken := ""
	username, err := utils.ExtractUsernameFromToken(emptyToken)

	// Assertions
	assert.Error(t, err, "Error should occur when extracting from an empty token")
	assert.Equal(t, "", username, "Username should be empty for an empty token")
	log.Println("Failed to extract username from empty token as expected")
}
