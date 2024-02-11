// jwt.go
package utils

import (
	"errors"
	"fmt"
	"main/internal/api/models"
	"main/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generates a JWT token with the provided username
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JWTKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the provided JWT token and returns the username if valid
func ValidateToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
