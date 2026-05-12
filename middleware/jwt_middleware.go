package middleware

import (
	"context"
	"net/http"
	"strings"

	"golang-api/config"
	"golang-api/logger"
	"golang-api/utils"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {

			logger.WarnLogger.Println(
				"Missing authorization token",
			)

			http.Error(
				w,
				"Missing token",
				http.StatusUnauthorized,
			)

			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {

			logger.WarnLogger.Println(
				"Invalid token format",
			)

			http.Error(
				w,
				"Invalid token format",
				http.StatusUnauthorized,
			)

			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				return utils.JwtSecret, nil
			},
		)

		if err != nil || !token.Valid {

			logger.WarnLogger.Println(
				"Invalid or expired token",
			)

			http.Error(
				w,
				"Invalid token",
				http.StatusUnauthorized,
			)

			return
		}

		// Check token in Redis
		_, err = config.RDB.Get(
			config.Ctx,
			tokenString,
		).Result()

		if err != nil {

			logger.WarnLogger.Println(
				"Token not found in Redis",
			)

			http.Error(
				w,
				"Invalid session",
				http.StatusUnauthorized,
			)

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {

			logger.ErrorLogger.Println(
				"Failed to parse JWT claims",
			)

			http.Error(
				w,
				"Invalid token claims",
				http.StatusUnauthorized,
			)

			return
		}

		userID, ok := claims["user_id"].(float64)

		if !ok {

			logger.ErrorLogger.Println(
				"Invalid user_id in token claims",
			)

			http.Error(
				w,
				"Invalid token payload",
				http.StatusUnauthorized,
			)

			return
		}

		logger.InfoLogger.Println(
			"Authenticated user:",
			int(userID),
		)

		logger.DebugLogger.Println(
			"JWT Claims:",
			claims,
		)

		ctx := context.WithValue(
			r.Context(),
			"userID",
			int(userID),
		)

		next(w, r.WithContext(ctx))
	}
}
