package utils

import (
	"database/sql"
	"fmt"
	"os"
	"user_post_creation/model"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func ConnectToDB(conn string) (*sql.DB, error) {
	var err error
	PSG, err := sql.Open("postgres", conn)
	if err != nil {
		log.Errorf("Unable to connect to postgres db %v", err)
		return nil, err
	}
	if err := PSG.Ping(); err != nil {
		return nil, err
	}
	return PSG, nil
}

func ValidateUser(username, password string) (bool, error) {
	var hashedPassword string
	PSG, err := ConnectToDB("")
	defer PSG.Close()
	if err != nil {
		log.Errorf("Unable to connect to DB %+v", err)
		return false, err
	}
	if err := PSG.QueryRow("SELECT password FROM users WHERE username =$1", username).Scan(&hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			log.Infof("user %v not found", username)
			return false, nil
		}
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Infof("Bcrypt password does not matched %+v", password)
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func RegisterUser(user model.User) error {
	if valid, err := ValidateUser(user.Email, user.Password); !valid || err != nil {
		return fmt.Errorf("user already present")
	}
	if user.ConfirmPassword != user.Password {
		return fmt.Errorf("confirm password did not matched")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	PSG, err := ConnectToDB("")
	if err != nil {
		log.Errorf("Unable to connect to DB %+v", err)
		return err
	}

	_, err = PSG.Exec("INSERT INTO users (username, password, firstname, lastname, mobile, id) VALUES ($1, $2, $3, $4, $5)", user.Email, string(hashedPassword), user.Fname, user.Lname, user.Id)
	PSG.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetUserDetails(userid string) (model.UserDetails, error) {
	PSG, err := ConnectToDB("")
	if err != nil {
		log.Errorf("Unable to connect to DB %+v", err)
		return model.UserDetails{}, err
	}
	var user model.User
	if err := PSG.QueryRow("SELECT username, firstname, lastname, mobile, id FROM users where id = $1", userid).Scan(&user.Email, &user.Fname, &user.Lname, &user.Mobile, &user.Id); err != nil {
		if err == sql.ErrNoRows {
			return model.UserDetails{}, err
		} else {
			return model.UserDetails{}, err
		}
	}

	var followersCount int
	if err := PSG.QueryRow("SELECT COUNT(*) FROM followers WHERE user_id = $1", userid).Scan(&followersCount); err != nil {
		return model.UserDetails{User: user, Followers: 0, ProfileImage: ""}, err
	}

	// Retrieve the profile image URL from the database
	var profileImageURL string
	if err = PSG.QueryRow("SELECT profile_image FROM users WHERE id = $1", userid).Scan(&profileImageURL); err != nil {
		return model.UserDetails{User: user, Followers: followersCount, ProfileImage: ""}, err
	}
	PSG.Close()
	return model.UserDetails{User: user, Followers: followersCount, ProfileImage: profileImageURL}, nil
}

func SavePost(post model.Post) (int64, error) {
	PSG, err := ConnectToDB("")
	if err != nil {
		log.Errorf("Unable to connect to DB %+v", err)
		return -1, err
	}
	result, err := PSG.Exec("INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3)", post.Authorid, post.Title, post.Content)
	if err != nil {
		return -1, err
	}
	PSG.Close()
	return result.LastInsertId()
}

func PostById(idStr string) (model.Post, error) {
	PSG, err := ConnectToDB("")
	if err != nil {
		return model.Post{}, err
	}
	var post model.Post
	if err := PSG.QueryRow("SELECT * FROM posts WHERE id = $1", idStr).Scan(&post.Id, &post.Authorid, &post.Title, &post.Content); err != nil {
		return model.Post{}, err
	}
	PSG.Close()
	return post, nil
}
