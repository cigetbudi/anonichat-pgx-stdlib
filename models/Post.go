package models

type Post struct {
	Id       int32  `json:"id"`
	Content  string `json:"content"`
	Location string `json:"location"`
	UserID   int64  `json:"user_id"`
}
