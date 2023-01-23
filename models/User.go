package models

type User struct {
	Id       uint32 `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	DOB      string `json:"DOB"`
}
