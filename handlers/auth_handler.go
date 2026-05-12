package handlers

import (
	"encoding/json"
	"net/http"

	"golang-api/config"
	"golang-api/logger"
	"golang-api/models"
	"golang-api/services"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to decode register request:",
			err,
		)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {

		logger.ErrorLogger.Println(
			"Password hashing failed:",
			err,
		)

		http.Error(
			w,
			"Failed to hash password",
			http.StatusInternalServerError,
		)

		return
	}

	query := `
	INSERT INTO users(name,email,password)
	VALUES(?,?,?)
	`

	result, err := config.DB.Exec(
		query,
		user.Name,
		user.Email,
		string(hashedPassword),
	)

	if err != nil {

		logger.ErrorLogger.Println(
			"User registration failed:",
			err,
		)

		http.Error(
			w,
			"Failed to register user",
			http.StatusInternalServerError,
		)

		return
	}

	id, err := result.LastInsertId()

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to get inserted user ID:",
			err,
		)

		http.Error(
			w,
			"Failed to process user registration",
			http.StatusInternalServerError,
		)

		return
	}

	user.ID = int(id)

	user.Password = ""

	logger.InfoLogger.Println(
		"User registered successfully:",
		user.Email,
	)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
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
