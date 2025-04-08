// Unit tests for password hashing and comparison
package tests

import (
	"testing"
	"backend/utils"
	"github.com/stretchr/testify/assert"
	"log"
)

// TestHashPassword_Success tests successful password hashing
func TestHashPassword_Success(t *testing.T) {
	password := "mySecurePassword"
	hash, err := utils.HashPassword(password)

	// Assertions
	assert.NoError(t, err, "Error should be nil when hashing password")
	assert.NotEmpty(t, hash, "Hash should not be empty")
	log.Println("Hash generated successfully:", hash)
}

// TestCheckPasswordHash_Success tests successful password comparison
func TestCheckPasswordHash_Success(t *testing.T) {
	password := "mySecurePassword"
	hash, _ := utils.HashPassword(password)

	// Compare hash and original password
	match := utils.CheckPasswordHash(password, hash)

	// Assertions
	assert.True(t, match, "Password should match the hash")
	log.Println("Password and hash match verified successfully")
}

// TestCheckPasswordHash_Failure tests failure when the password does not match the hash
func TestCheckPasswordHash_Failure(t *testing.T) {
	password := "mySecurePassword"
	hash, _ := utils.HashPassword(password)

	// Use a different password for comparison
	wrongPassword := "wrongPassword"
	match := utils.CheckPasswordHash(wrongPassword, hash)

	// Assertions
	assert.False(t, match, "Password should not match the hash")
	log.Println("Password and hash do not match as expected")
}

// TestHashPassword_Empty tests hashing an empty password
func TestHashPassword_Empty(t *testing.T) {
	password := ""
	hash, err := utils.HashPassword(password)

	// Assertions
	assert.NoError(t, err, "Error should not occur even with an empty password")
	assert.NotEmpty(t, hash, "Hash should not be empty for an empty password")
	log.Println("Empty password hashing handled successfully")
}
