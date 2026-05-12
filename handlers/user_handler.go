package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-api/logger"
	"golang-api/models"
	"golang-api/services"
	"golang-api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to decode create user request:",
			err,
		)

		utils.SendError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	id, err := services.CreateUserService(user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Create user failed:",
			err,
		)

		utils.SendError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	user.ID = int(id)

	logger.InfoLogger.Println(
		"User created successfully:",
		user.Email,
	)

	utils.SendSuccess(
		w,
		http.StatusCreated,
		"User created successfully",
		user,
	)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := services.GetUsersService()

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to fetch users:",
			err,
		)

		utils.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to fetch users",
		)

		return
	}

	logger.InfoLogger.Println(
		"Fetched all users successfully",
	)

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Users fetched successfully",
		users,
	)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	err := services.DeleteUserService(id)

	if err != nil {

		logger.ErrorLogger.Println(
			"Delete user failed:",
			err,
		)

		utils.SendError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	logger.InfoLogger.Println(
		"User deleted successfully:",
		id,
	)

	utils.SendSuccess(
		w,
		http.StatusOK,
		"User deleted successfully",
		nil,
	)
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

		utils.SendError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	err = services.UpdateUserService(user, id)

	if err != nil {

		logger.ErrorLogger.Println(
			"Update user failed:",
			err,
		)

		utils.SendError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	user.ID = atoi(id)

	logger.InfoLogger.Println(
		"User updated successfully:",
		id,
	)

	utils.SendSuccess(
		w,
		http.StatusOK,
		"User updated successfully",
		user,
	)
}
