package utils

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(password, hash string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash)); err != nil {
		return errors.New("invalid password")
	}

	return nil
}
