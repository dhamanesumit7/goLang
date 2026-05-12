package repository

import (
	"golang-api/config"
	"golang-api/models"
)

func GetUserByEmail(email string) (models.User, error) {

	var user models.User

	query := `
	SELECT id,name,email,password
	FROM users
	WHERE email=?
	`

	err := config.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	return user, err
}