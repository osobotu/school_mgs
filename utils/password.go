package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %s", err)
	}
	return string(hash), nil
}

func CheckPassword(password, password_hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
}
