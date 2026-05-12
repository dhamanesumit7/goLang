package services

import (
	"errors"
	"time"

	"golang-api/config"
	"golang-api/logger"
	"golang-api/models"
	"golang-api/repository"
	"golang-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func LoginService(input models.User) (string, error) {

	user, err := repository.GetUserByEmail(input.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	tokenString, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return "", err
	}

	// Store token in Redis
	err = config.RDB.Set(
		config.Ctx,
		tokenString,
		user.ID,
		time.Minute*10,
	).Err()

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to store token in Redis:",
			err,
		)

		return "", err
	}

	logger.InfoLogger.Println(
		"Token stored in Redis successfully",
	)

	return tokenString, nil
}
