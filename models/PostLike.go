package models

import (
	"time"

	"github.com/google/uuid"
)

type PostLike struct {
	Id        uuid.UUID `json:"id"`
	PostId    uuid.UUID `json:"post_id"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CountLike struct {
	Likes int32 `json:"likes"`
}
