package services

import (
	"errors"
	"strings"

	"golang-api/logger"
	"golang-api/models"
	"golang-api/repository"
	"golang-api/utils"
)

func CreateUserService(
	user models.User,
) (int64, error) {

	validationErrors := utils.ValidateUser(user)

	if len(validationErrors) > 0 {

		// logger.WarnLogger.Println(
		// 	"User validation failed:",
		// 	validationErrors,
		// )

		return 0, errors.New(
			strings.Join(validationErrors, ", "),
		)
	}

	id, err := repository.CreateUserRepo(user)

	if err != nil {

		// logger.ErrorLogger.Println(
		// 	"Create user repository error:",
		// 	err,
		// )

		return 0, err
	}

	logger.InfoLogger.Println(
		"CreateUserService executed successfully",
	)

	return id, nil
}

func GetUsersService() ([]models.User, error) {

	users, err := repository.GetUsersRepo()

	if err != nil {

		// logger.ErrorLogger.Println(
		// 	"Get users repository error:",
		// 	err,
		// )

		return nil, err
	}

	// logger.InfoLogger.Println(
	// 	"GetUsersService executed successfully",
	// )

	return users, nil
}

func UpdateUserService(
	user models.User,
	id string,
) error {

	validationErrors := utils.ValidateUser(user)

	if len(validationErrors) > 0 {

		// logger.WarnLogger.Println(
		// 	"Update validation failed:",
		// 	validationErrors,
		// )

		return errors.New(
			strings.Join(validationErrors, ", "),
		)
	}

	err := repository.UpdateUserRepo(user, id)

	if err != nil {

		// logger.ErrorLogger.Println(
		// 	"Update repository error:",
		// 	err,
		// )

		return err
	}

	logger.InfoLogger.Println(
		"UpdateUserService executed successfully",
	)

	return nil
}

func DeleteUserService(id string) error {

	err := repository.DeleteUserRepo(id)

	if err != nil {

		// logger.ErrorLogger.Println(
		// 	"Delete repository error:",
		// 	err,
		// )

		return err
	}

	// logger.InfoLogger.Println(
	// 	"DeleteUserService executed successfully",
	// )

	return nil
}
