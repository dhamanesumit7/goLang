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

	err := utils.ValidateLogin(input)

	if err != nil {

		logger.WarnLogger.Println(
			"Login validation failed:",
			err,
		)

		return "", err
	}

	user, err := repository.GetUserByEmail(
		input.Email,
	)

	if err != nil {

		logger.WarnLogger.Println(
			"Invalid login email:",
			input.Email,
		)

		return "", errors.New(
			"invalid email or password",
		)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {

		logger.WarnLogger.Println(
			"Invalid password attempt:",
			input.Email,
		)

		return "", errors.New(
			"invalid email or password",
		)
	}

	tokenString, err := utils.GenerateJWT(user.ID)

	if err != nil {

		logger.ErrorLogger.Println(
			"JWT generation failed:",
			err,
		)

		return "", err
	}

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

	logger.InfoLogger.Println(
		"User login successful:",
		input.Email,
	)

	return tokenString, nil
}

func LogoutService(tokenString string) error {

	err := config.RDB.Del(
		config.Ctx,
		tokenString,
	).Err()

	if err != nil {

		logger.ErrorLogger.Println(
			"Failed to delete token from Redis:",
			err,
		)

		return err
	}

	logger.InfoLogger.Println(
		"User logged out successfully",
	)

	return nil
}

func RegisterService(user models.User) (models.User, error) {

	err := utils.ValidateUser(user)

	if err != nil {

		logger.WarnLogger.Println(
			"Register validation failed:",
			err,
		)

		return models.User{}, err
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

		return models.User{}, err
	}

	user.Password = string(hashedPassword)

	id, err := repository.CreateUserRepo(user)

	if err != nil {

		logger.ErrorLogger.Println(
			"Register repository error:",
			err,
		)

		return models.User{}, err
	}

	user.ID = int(id)

	user.Password = ""

	logger.InfoLogger.Println(
		"User registered successfully:",
		user.Email,
	)

	return user, nil
}
