package models

import "time"

type PostLike struct {
	Id        int32     `json:"id"`
	PostId    int32     `json:"post_id"`
	UserId    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
