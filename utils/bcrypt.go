package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

// GenerateHashFromPassword takes a plain password and returns the bcrypt hash
func GenerateHash(password string) (string, error) {
	// bcrypt.DefaultCost is a reasonable default cost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// MatchPassword compares a bcrypt hashed password with a plain password
func MatchPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
