package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"golang-api/logger"
	"golang-api/models"
	"golang-api/services"
	"golang-api/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to decode register request:",
			err,
		)

		utils.SendError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	createdUser, err := services.RegisterService(user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Register service failed:",
			err,
		)

		utils.SendError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	utils.SendSuccess(
		w,
		http.StatusCreated,
		"User registered successfully",
		createdUser,
	)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var input models.User

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to decode login request:",
			err,
		)

		utils.SendError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	tokenString, err := services.LoginService(input)

	if err != nil {

		logger.WarnLogger.Println(
			"Login failed:",
			input.Email,
		)

		utils.SendError(
			w,
			http.StatusUnauthorized,
			err.Error(),
		)

		return
	}

	logger.InfoLogger.Println(
		"Login successful:",
		input.Email,
	)

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Login successful",
		map[string]string{
			"token": tokenString,
		},
	)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {

		logger.WarnLogger.Println(
			"Missing authorization token during logout",
		)

		utils.SendError(
			w,
			http.StatusUnauthorized,
			"Missing token",
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

		utils.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to logout",
		)

		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Logout successful",
		nil,
	)
}
