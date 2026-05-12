package handlers

import (
	"encoding/json"
	"net/http"

	"golang-api/config"
	"golang-api/logger"
	"golang-api/models"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userID").(int)	

	logger.InfoLogger.Println(
		"Profile access request:",
		userID,
	)

	var user models.User

	query := `
	SELECT id,name,email
	FROM users
	WHERE id=?
	`

	err := config.DB.QueryRow(query, userID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {

		logger.WarnLogger.Println(
			"Profile not found for user ID:",
			userID,
		)

		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)

		return
	}

	logger.InfoLogger.Println(
		"Profile fetched successfully:",
		user.Email,
	)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
