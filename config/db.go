package config

import (
	"database/sql"
	"fmt"
	"os"

	"golang-api/logger"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {

	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	DB, err = sql.Open("mysql", dsn)

	if err != nil {

		logger.ErrorLogger.Println(
			"Database open error:",
			err,
		)

		panic(err)
	}

	err = DB.Ping()

	if err != nil {

		logger.ErrorLogger.Println(
			"Database connection failed:",
			err,
		)

		panic(err)
	}

	logger.InfoLogger.Println(
		"MySQL connected successfully",
	)
}
