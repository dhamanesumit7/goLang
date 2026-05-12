package main

import (
	"net/http"

	"golang-api/config"
	"golang-api/logger"
	"golang-api/routes"
)

func main() {

	config.LoadEnv()

	config.ConnectDB()

	config.ConnectRedis()

	routes.SetupRoutes()

	logger.InfoLogger.Println(
		"Server started on port 8080",
	)

	http.ListenAndServe(":8080", nil)
}
