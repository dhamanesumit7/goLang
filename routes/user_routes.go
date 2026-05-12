package routes

import (
	"net/http"

	"golang-api/handlers"
	"golang-api/middleware"
)

func UserRoutes() {

	http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case "GET":
			handlers.GetUsers(w, r)

		case "POST":
			handlers.CreateUser(w, r)

		case "PUT":
			handlers.UpdateUser(w, r)

		case "DELETE":
			handlers.DeleteUser(w, r)

		default:
			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)
		}
	})

	http.HandleFunc(
		"/api/v1/users/me",
		middleware.JWTAuth(handlers.Profile),
	)
}
