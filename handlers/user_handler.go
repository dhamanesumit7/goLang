package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-api/config"
	"golang-api/logger"
	"golang-api/models"
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

	query := `
	INSERT INTO users(name,email,password)
	VALUES(?,?,?)
	`

	result, err := config.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {

		logger.ErrorLogger.Println(
			"Create user database error:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	id, err := result.LastInsertId()

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to fetch inserted user ID:",
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

	rows, err := config.DB.Query(
		"SELECT id,name,email FROM users",
	)

	if err != nil {

		logger.ErrorLogger.Println(
			"Get users database error:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)

		if err != nil {

			logger.ErrorLogger.Println(
				"Row scan error:",
				err,
			)

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)

			return
		}

		users = append(users, user)
	}

	logger.InfoLogger.Println(
		"Fetched all users successfully",
	)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	query := "DELETE FROM users WHERE id=?"

	result, err := config.DB.Exec(query, id)

	if err != nil {

		logger.ErrorLogger.Println(
			"Delete user database error:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to get rows affected:",
			err,
		)

		http.Error(
			w,
			"Failed to delete user",
			http.StatusInternalServerError,
		)

		return
	}

	if rowsAffected == 0 {

		logger.WarnLogger.Println(
			"Delete failed - user not found:",
			id,
		)

		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
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

	query := `
	UPDATE users
	SET name=?, email=?, password=?
	WHERE id=?
	`

	result, err := config.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		id,
	)

	if err != nil {

		logger.ErrorLogger.Println(
			"Update user database error:",
			err,
		)

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to get rows affected:",
			err,
		)

		http.Error(
			w,
			"Failed to update user",
			http.StatusInternalServerError,
		)

		return
	}

	if rowsAffected == 0 {

		logger.WarnLogger.Println(
			"Update failed - user not found:",
			id,
		)

		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
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
