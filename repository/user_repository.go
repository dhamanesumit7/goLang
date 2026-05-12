package repository

import (
	"fmt"

	"golang-api/config"
	"golang-api/models"
)

func CreateUserRepo(user models.User) (int64, error) {

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
		return 0, err
	}

	return result.LastInsertId()
}

func GetUsersRepo() ([]models.User, error) {

	rows, err := config.DB.Query(
		"SELECT id,name,email FROM users",
	)

	if err != nil {
		return nil, err
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
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateUserRepo(user models.User, id string) error {

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
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func DeleteUserRepo(id string) error {

	query := "DELETE FROM users WHERE id=?"

	result, err := config.DB.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

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
