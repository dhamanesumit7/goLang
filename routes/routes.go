package routes

import "golang-api/logger"

func SetupRoutes() {

	logger.InfoLogger.Println(
		"Initializing application routes",
	)

	AuthRoutes()

	UserRoutes()

	logger.InfoLogger.Println(
		"All routes initialized successfully",
	)
}
