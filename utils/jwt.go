package utils

import (
    "github.com/golang-jwt/jwt"
    "time"
    "log"
	"errors"
)
var jwtKey = []byte("secret")

// CustomClaims structure to include username
type CustomClaims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
    log.Println("Generating JWT for:", username)

    claims := &CustomClaims{
        Username: username,  // Set username field correctly
        StandardClaims: jwt.StandardClaims{
            Subject:   username,  // Set "sub" field correctly
            ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        log.Println("Error generating token:", err)
        return "", err
    }
    log.Println("Generated JWT:", tokenString)
    return tokenString, nil
}

// Generate Refresh Token
func GenerateRefreshToken(username string) (string, error) {
    claims := &CustomClaims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(1 * 24 * time.Hour).Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// Extract Username from JWT
func ExtractUsernameFromToken(tokenString string) (string, error) {
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        return "", errors.New("invalid token")
    }

    claims, ok := token.Claims.(*CustomClaims)
    if !ok || claims.Username == "" {
        return "", errors.New("invalid token claims")
    }

    return claims.Username, nil
}

