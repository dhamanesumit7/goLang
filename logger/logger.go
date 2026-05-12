package logger

import (
	"log"
	"os"
)

var (
	InfoLogger = log.New(
		os.Stdout,
		"[INFO] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	WarnLogger = log.New(
		os.Stdout,
		"[WARN] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	ErrorLogger = log.New(
		os.Stdout,
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	DebugLogger = log.New(
		os.Stdout,
		"[DEBUG] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
)
