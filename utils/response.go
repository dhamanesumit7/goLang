package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func SendSuccess(
	w http.ResponseWriter,
	statusCode int,
	message string,
	data interface{},
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(
		SuccessResponse{
			Success: true,
			Message: message,
			Data:    data,
		},
	)
}

func SendError(
	w http.ResponseWriter,
	statusCode int,
	errorMessage string,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(
		ErrorResponse{
			Success: false,
			Error:   errorMessage,
		},
	)
}
