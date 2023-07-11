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
	CreateAt        time.Time `json:"createdat"`
	Udate           time.Time `json:"updatedat`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Msg    string      `json:"error"`
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type Config struct {
	AwsRegion    string `json:"awsregion"`
	AwsID        string `json:"awsid"`
	AwsSecretKey string `json:"awssecret"`
	PSGUser      string `json:"psguser"`
	PSGHost      string `json:"psghost"`
	PSGPass      string `json:"psgpass"`
	PSGDB        string `json:"psgdb"`
	PSGPort      int    `json:"psgport"`
	S3Bucket     string `json:"s3bucket"`
	ServiceHost  string `json:"servicehost"`
}
