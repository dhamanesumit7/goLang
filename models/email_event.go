package models

type EmailEvent struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Event   string `json:"event"`
	Message string `json:"message"`
}
