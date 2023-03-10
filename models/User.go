package models

import "github.com/google/uuid"

type User struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Fullname   string    `json:"fullname"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	DOB        string    `json:"DOB"`
	Phone      string    `json:"phone"`
	GenderCode string    `json:"gender_code"`
}

// type User struct {
// 	Id       uint32 `json:"id"`
// 	Username string `json:"username"`
// 	Fullname string `json:"fullname"`
// 	Password string `json:"password"`
// 	Email    string `json:"email"`
// 	DOB      string `json:"DOB"`
// }
