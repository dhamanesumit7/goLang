package utils

import (
	"errors"
	"regexp"

	"golang-api/models"
)

func ValidateUser(user models.User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if len(user.Password) < 6 {
		return errors.New(
			"password must be at least 6 characters",
		)
	}

	return nil
}

func ValidateLogin(user models.User) error {

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

func isValidEmail(email string) bool {

	emailRegex := regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
	)

	return emailRegex.MatchString(email)
}
