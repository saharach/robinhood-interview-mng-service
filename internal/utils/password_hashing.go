package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a salt, hashes the password with the salt using bcrypt,
// and returns the base64-encoded salt, hashed password, and error.
func HashPassword(password string) (string, string, error) {
	// Generate a random salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", err
	}
	passwordWithSalt := append([]byte(password), salt...)

	// Hash the password with the salt using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, 16)
	if err != nil {
		return "", "", err
	}

	// Encode the salt to base64
	saltBase64 := base64.StdEncoding.EncodeToString(salt)

	// Encode the hashed password to base64
	hashedPasswordBase64 := base64.StdEncoding.EncodeToString(hashedPassword)

	return saltBase64, hashedPasswordBase64, nil
}

// ComparePasswordHash compares a password with a hashed password and salt using bcrypt.
func ComparePasswordHash(password, hashedPasswordBase64, saltBase64 string) error {
	// Decode the salt from base64
	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return err
	}

	// Decode the hashed password from base64
	hashedPassword, err := base64.StdEncoding.DecodeString(hashedPasswordBase64)
	if err != nil {
		return err
	}

	// Combine the provided password with the salt
	passwordWithSalt := append([]byte(password), salt...)

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword(hashedPassword, passwordWithSalt); err != nil {
		return errors.New("Not match")
	}

	return nil
}
