package services

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"golang-api/config"
	"golang-api/logger"
	"golang-api/models"
	"golang-api/repository"
	"golang-api/utils"

	"github.com/segmentio/kafka-go"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(input models.User) (string, error) {

	validationErrors := utils.ValidateLogin(input)

	if len(validationErrors) > 0 {

		// logger.WarnLogger.Println(
		// 	"Login validation failed:",
		// 	validationErrors,
		// )

		return "", errors.New(
			strings.Join(validationErrors, ", "),
		)
	}

	user, err := repository.GetUserByEmail(
		input.Email,
	)

	if err != nil {

		// logger.WarnLogger.Println(
		// 	"Invalid login email:",
		// 	input.Email,
		// )

		return "", errors.New(
			"invalid email or password",
		)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {

		// logger.WarnLogger.Println(
		// 	"Invalid password attempt:",
		// 	input.Email,
		// )

		return "", errors.New(
			"invalid email or password",
		)
	}

	tokenString, err := utils.GenerateJWT(user.ID)

	if err != nil {

		// logger.ErrorLogger.Println(
		// 	"JWT generation failed:",
		// 	err,
		// )

		return "", err
	}

	err = config.RDB.Set(
		config.Ctx,
		tokenString,
		user.ID,
		time.Minute*10,
	).Err()

	if err != nil {

		// logger.ErrorLogger.Println(
		// 	"Failed to store token in Redis:",
		// 	err,
		// )

		return "", err
	}

	logger.InfoLogger.Println(
		"Token stored in Redis successfully",
	)

	// logger.InfoLogger.Println(
	// 	"User login successful:",
	// 	input.Email,
	// )

	loginEvent := models.LoginEvent{
		UserID: user.ID,
		Email:  user.Email,
		Event:  "USER_LOGIN",
	}

	messageBytes, err := json.Marshal(
		loginEvent,
	)

	if err == nil {

		message := kafka.Message{
			Value: messageBytes,
		}

		err = config.LoginWriter.WriteMessages(
			context.Background(),
			message,
		)

		if err != nil {

			logger.ErrorLogger.Println(
				"Kafka login event publish failed:",
				err,
			)

		} else {

			logger.InfoLogger.Println(
				"Login event published to Kafka",
			)
		}
	}

	return tokenString, nil
}

func LogoutService(tokenString string) error {

	err := config.RDB.Del(
		config.Ctx,
		tokenString,
	).Err()

	if err != nil {

		// logger.ErrorLogger.Println(
		// 	"Failed to delete token from Redis:",
		// 	err,
		// )

		return err
	}

	logger.InfoLogger.Println(
		"User logged out successfully",
	)

	return nil
}

func RegisterService(
	user models.User,
) (models.User, error) {

	// logger.DebugLogger.Println(
	// 	"Register API hit",
	// )

	validationErrors := utils.ValidateUser(user)

	if len(validationErrors) > 0 {

		return models.User{}, errors.New(
			strings.Join(validationErrors, ", "),
		)
	}

	// Check existing user BEFORE insert
	exists, err := repository.CheckUserExists(
		user.Email,
	)

	if err != nil {

		return models.User{}, err
	}

	if exists {

		return models.User{}, errors.New(
			"user already exists, please login",
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {

		return models.User{}, err
	}

	user.Password = string(hashedPassword)

	id, err := repository.CreateUserRepo(user)

	if err != nil {

		return models.User{}, err
	}

	user.ID = int(id)

	user.Password = ""

	logger.InfoLogger.Println(
		"User registered successfully:",
		user.Email,
	)

	// logger.DebugLogger.Println(
	// 	"Register request received at:",
	// 	time.Now(),
	// )

	// Kafka Email Event
	emailEvent := models.EmailEvent{
		Email: user.Email,

		Name: user.Name,

		Event: "SEND_WELCOME_EMAIL",

		Message: "Welcome to Dev community",
	}

	messageBytes, err := json.Marshal(
		emailEvent,
	)

	if err == nil {

		message := kafka.Message{
			Value: messageBytes,
		}

		err = config.EmailWriter.WriteMessages(
			context.Background(),
			message,
		)

		if err != nil {

			logger.ErrorLogger.Println(
				"Kafka email event publish failed:",
				err,
			)

		} else {

			logger.InfoLogger.Println(
				"Email event published to Kafka",
			)
		}
	}

	return user, nil
}
