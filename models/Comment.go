package models

import "github.com/google/uuid"

type Comment struct {
	Id       uuid.UUID `json:"id"`
	PostId   uuid.UUID `json:"post_id"`
	UserId   uuid.UUID `json:"user_id"`
	Comment  string    `json:"comment"`
	Location string    `json:"location"`
}
