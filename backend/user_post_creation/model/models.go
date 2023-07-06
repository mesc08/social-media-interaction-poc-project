package model

import "time"

type User struct {
	Id              string    `json:"id,omitempty"`
	Fname           string    `json:"firstname"`
	Lname           string    `json:"lastname"`
	Email           string    `json:"email"`
	Mobile          int64     `json:"mobile"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirmpassword,omitempty"`
	ProfileImage    string    `json:"profile_image"`
	CreateAt        time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Msg    string      `json:"error"`
}
