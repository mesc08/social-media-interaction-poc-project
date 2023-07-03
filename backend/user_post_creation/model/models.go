package model

type User struct {
	Id              string `json:"id,omitempty"`
	Fname           string `json:"firstname"`
	Lname           string `json:"lastname"`
	Email           string `json:"email"`
	Mobile          int64  `json:"mobile"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword,omitempty"`
}

type Post struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	Authorid string `json:"authorid"`
	Likes    int    `json:"Likes"`
}

type Comment struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	Authorid string `json:"authorid"`
	Postid   string `json:"postid"`
	Likes    int    `json:"Likes"`
}
