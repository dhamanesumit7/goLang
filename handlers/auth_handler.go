package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"golang-api/logger"
	"golang-api/models"
	"golang-api/services"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to decode register request:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	createdUser, err := services.RegisterService(user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Register service failed:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(createdUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var input models.User

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to decode login request:",
			err,
		)

		http.Error(
			w,
			"Invalid request payload",
			http.StatusBadRequest,
		)

		return
	}

	tokenString, err := services.LoginService(input)

	if err != nil {

		logger.WarnLogger.Println(
			"Login failed for user:",
			input.Email,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusUnauthorized,
		)

		return
	}

	logger.InfoLogger.Println(
		"Login successful:",
		input.Email,
	)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {

		logger.WarnLogger.Println(
			"Missing authorization token during logout",
		)

		http.Error(
			w,
			"Missing token",
			http.StatusUnauthorized,
		)

		return
	}

	tokenString := strings.TrimPrefix(
		authHeader,
		"Bearer ",
	)

	err := services.LogoutService(tokenString)

	if err != nil {

		logger.ErrorLogger.Println(
			"Logout failed:",
			err,
		)

		http.Error(
			w,
			"Failed to logout",
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}
