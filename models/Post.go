package models

import "github.com/google/uuid"

type Post struct {
	Id       uuid.UUID `json:"id"`
	Content  string    `json:"content"`
	Location string    `json:"location"`
	UserID   uuid.UUID `json:"user_id"`
}
