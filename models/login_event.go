package models

type LoginEvent struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Event  string `json:"event"`
}
