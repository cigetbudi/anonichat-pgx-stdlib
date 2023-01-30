package models

type Comment struct {
	Id       int32  `json:"id"`
	PostId   int64  `json:"post_id"`
	UserId   int64  `json:"user_id"`
	Comment  string `json:"comment"`
	Location string `json:"location"`
}
