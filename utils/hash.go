package utils

import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

// HashPassword hashes a plain-text password
func HashPassword(password string) (string, error) {
    log.Println("Hashing password:", password)
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Error while hashing password:", err)
        return "", err
    }
    log.Println("Generated hash:", string(bytes))
    return string(bytes), nil
}

// CheckPasswordHash compares a hashed password with a plain-text one
func CheckPasswordHash(password, hash string) bool {
    log.Println("Comparing password and hash")
    log.Println("Plain Password:", password)
    log.Println("Stored Hash:", hash)

    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        log.Println("Password comparison failed:", err)
        return false
    }
    log.Println("Password comparison successful")
    return true
}
