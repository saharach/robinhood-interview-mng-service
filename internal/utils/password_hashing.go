// utils/password_hashing.go

package utils

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a salt, hashes the password with the salt using bcrypt,
// and returns the hashed password and salt.
func HashPassword(password string) ([]byte, []byte, error) {
	// Generate a random salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, nil, err
	}

	// Hash the password with the salt using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	return hashedPassword, salt, nil
}

// ComparePasswordHash compares a password with a hashed password and salt using bcrypt.
func ComparePasswordHash(password string, hashedPassword, salt []byte) error {
	// Combine the password with the salt
	passwordWithSalt := append([]byte(password), salt...)

	// Hash the combined password and salt
	newHashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Compare the newly hashed password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword(hashedPassword, newHashedPassword); err != nil {
		return err
	}

	return nil
}
