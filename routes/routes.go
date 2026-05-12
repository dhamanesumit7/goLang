package routes

import (
	"net/http"

	"golang-api/handlers"
	"golang-api/logger"
	"golang-api/middleware"
)

func SetupRoutes() {

	logger.InfoLogger.Println(
		"Initializing application routes",
	)

	http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {

		logger.InfoLogger.Println(
			"Incoming request:",
			r.Method,
			r.URL.Path,
		)

		switch r.Method {

		case "GET":

			logger.InfoLogger.Println(
				"GET /api/v1/users route accessed",
			)

			handlers.GetUsers(w, r)

		case "POST":

			logger.InfoLogger.Println(
				"POST /api/v1/users route accessed",
			)

			handlers.CreateUser(w, r)

		case "PUT":

			logger.InfoLogger.Println(
				"PUT /api/v1/users route accessed",
			)

			handlers.UpdateUser(w, r)

		case "DELETE":

			logger.InfoLogger.Println(
				"DELETE /api/v1/users route accessed",
			)

			handlers.DeleteUser(w, r)

		default:

			logger.WarnLogger.Println(
				"Method not allowed:",
				r.Method,
			)

			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)
		}
	})

	logger.InfoLogger.Println(
		"Public route enabled: /api/v1/auth/register",
	)

	http.HandleFunc(
		"/api/v1/auth/register",
		handlers.RegisterUser,
	)

	logger.InfoLogger.Println(
		"Public route enabled: /api/v1/auth/login",
	)

	http.HandleFunc(
		"/api/v1/auth/login",
		handlers.LoginUser,
	)

	logger.InfoLogger.Println(
		"Protected route enabled: /api/v1/users/me",
	)

	http.HandleFunc(
		"/api/v1/users/me",
		middleware.JWTAuth(handlers.Profile),
	)
}
