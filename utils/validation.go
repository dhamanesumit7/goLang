package utils

import (
	"regexp"

	"golang-api/models"
)

func ValidateUser(user models.User) []string {

	var errors []string

	if user.Name == "" {

		errors = append(
			errors,
			"name is required",
		)
	}

	if user.Email == "" {

		errors = append(
			errors,
			"email is required",
		)

	} else if !isValidEmail(user.Email) {

		errors = append(
			errors,
			"invalid email format",
		)
	}

	if user.Password == "" {

		errors = append(
			errors,
			"password is required",
		)

	} else if len(user.Password) < 6 {

		errors = append(
			errors,
			"password must be at least 6 characters",
		)
	}

	return errors
}

func ValidateLogin(user models.User) []string {

	var errors []string

	if user.Email == "" {

		errors = append(
			errors,
			"email is required",
		)
	}

	if user.Password == "" {

		errors = append(
			errors,
			"password is required",
		)
	}

	return errors
}

func isValidEmail(email string) bool {

	emailRegex := regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
	)

	return emailRegex.MatchString(email)
}
