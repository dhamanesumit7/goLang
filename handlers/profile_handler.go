package handlers

import (
	"net/http"

	"golang-api/config"
	"golang-api/logger"
	"golang-api/models"
	"golang-api/utils"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userID").(int)

	logger.InfoLogger.Println(
		"Profile accessed:",
		userID,
	)

	var user models.User

	query := `
	SELECT id,name,email
	FROM users
	WHERE id=?
	`

	err := config.DB.QueryRow(
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {

		logger.ErrorLogger.Println(
			"Profile fetch failed:",
			err,
		)

		utils.SendError(
			w,
			http.StatusNotFound,
			"User not found",
		)

		return
	}

	utils.SendSuccess(
		w,
		http.StatusOK,
		"Profile fetched successfully",
		user,
	)
}
