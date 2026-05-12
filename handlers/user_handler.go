package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-api/logger"
	"golang-api/models"
	"golang-api/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to decode create user request:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	id, err := services.CreateUserService(user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Create user failed:",
			err,
		)

		http.Error(
			w,
			"Failed to create user",
			http.StatusInternalServerError,
		)

		return
	}

	user.ID = int(id)

	logger.InfoLogger.Println(
		"User created successfully:",
		user.Email,
	)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := services.GetUsersService()

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to fetch users:",
			err,
		)

		http.Error(
			w,
			"Failed to fetch users",
			http.StatusInternalServerError,
		)

		return
	}

	logger.InfoLogger.Println(
		"Fetched all users successfully",
	)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	err := services.DeleteUserService(id)

	if err != nil {

		logger.ErrorLogger.Println(
			"Delete user failed:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	logger.InfoLogger.Println(
		"User deleted successfully:",
		id,
	)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}

func atoi(s string) int {

	var num int

	fmt.Sscanf(s, "%d", &num)

	return num
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to decode update request:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	err = services.UpdateUserService(user, id)

	if err != nil {

		logger.ErrorLogger.Println(
			"Update user failed:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	user.ID = atoi(id)

	logger.InfoLogger.Println(
		"User updated successfully:",
		id,
	)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

