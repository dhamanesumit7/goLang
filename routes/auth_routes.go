package routes

import (
	"net/http"

	"golang-api/handlers"
)

func AuthRoutes() {

	http.HandleFunc(
		"/api/v1/auth/register",
		handlers.RegisterUser,
	)

	http.HandleFunc(
		"/api/v1/auth/login",
		handlers.LoginUser,
	)

	http.HandleFunc(
		"/api/v1/auth/logout",
		handlers.LogoutUser,
	)
}
