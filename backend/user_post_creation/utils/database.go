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
