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
	Id       int64  `json:"id"`
	Content  string `json:"content"`
	Authorid string `json:"authorid"`
	Likes    int    `json:"Likes"`
	Title    string `json:"title"`
}

type Comment struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	Authorid string `json:"authorid"`
	Postid   string `json:"postid"`
	Likes    int    `json:"Likes"`
}

type UserDetails struct {
	User         User   `json:"user"`
	Followers    int    `json:"followers"`
	ProfileImage string `json:"profile_image"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Msg    string      `json:"error"`
}
