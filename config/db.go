package config

import (
	"database/sql"
	"golang-api/logger"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {

	var err error

	DB, err = sql.Open(
		"mysql",
		"root:S@um9594@tcp(127.0.0.1:3306)/golangdb",
	)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()

	if err != nil {

		logger.ErrorLogger.Println("Database connection failed:", err)

		panic(err)
	}

	logger.InfoLogger.Println("MySQL connected successfully")
}
