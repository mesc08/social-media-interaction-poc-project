package utils

import (
	"database/sql"
	"fmt"
	"os"
	"user_post_creation/config"
	"user_post_creation/model"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func GetPSGConn() string {

	cfg := model.DBConfig{
		User:     config.ViperConfig.PSGUser,
		Password: config.ViperConfig.PSGPass,
		Host:     config.ViperConfig.PSGHost,
		Port:     config.ViperConfig.PSGPort,
		Database: config.ViperConfig.PSGDB,
	}
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%v sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port)
}

func ConnectToDB() (*sql.DB, error) {
	conn := GetPSGConn()
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
	PSG, err := ConnectToDB()
	if err != nil {
		log.Errorf("Unable to connect to DB %+v", err)
		return false, err
	}
	defer PSG.Close()
	createTableQuery := CreateUserTableIfNotExist()
	if _, err := PSG.Exec(createTableQuery); err != nil {
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
	PSG, err := ConnectToDB()
	if err != nil {
		log.Errorf("Unable to connect to DB %+v", err)
		return err
	}

	createTableQuery := CreateUserTableIfNotExist()

	_, err = PSG.Exec(createTableQuery)
	if err != nil {
		return err
	}
	_, err = PSG.Exec("INSERT INTO users (username, password, firstname, lastname, mobile, userid, profilepic) VALUES ($1, $2, $3, $4, $5, $6)", user.Email, string(hashedPassword), user.Fname, user.Lname, user.Id, user.ProfileImage)
	PSG.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetUserDetails(userid string) (model.User, error) {
	PSG, err := ConnectToDB()
	if err != nil {
		log.Errorf("Unable to connect to DB %+v", err)
		return model.User{}, err
	}
	createTableQuery := CreateUserTableIfNotExist()

	_, err = PSG.Exec(createTableQuery)
	if err != nil {
		return model.User{}, err
	}
	var user model.User
	if err := PSG.QueryRow("SELECT username, firstname, lastname, mobile, userid, profilepic FROM users where userid = $1", userid).Scan(&user.Email, &user.Fname, &user.Lname, &user.Mobile, &user.Id); err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, err
		} else {
			return model.User{}, err
		}
	}
	PSG.Close()
	return user, nil
}

func DeleteUSer(userid string) error {
	PSG, err := ConnectToDB()
	if err != nil {
		return err
	}
	defer PSG.Close()
	stmt, err := PSG.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userid)
	if err != nil {
		return err
	}
	return nil

}

func EditUserDetails(userID string, updatedUser model.User) error {
	PSG, err := ConnectToDB()
	if err != nil {
		return err
	}
	defer PSG.Close()
	stmt, err := PSG.Prepare("UPDATE user SET name = $1, email = $2 WHERE id = $3")

	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(updatedUser.Fname, updatedUser.Email, updatedUser.Id); err != nil {
		return err
	}
	return nil
}

func StoreFileInUserDetails(fileURl string, id string) error {
	PSG, err := ConnectToDB()
	if err != nil {
		return err
	}
	defer PSG.Close()
	stmt, err := PSG.Prepare("UPDATE user SET profilepic = $1 WHERE id = $2")

	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(fileURl, id); err != nil {
		return err
	}
	return nil
}

func CreateUserTableIfNotExist() string {
	return `
	CREATE TABLE IF NOT EXISTS users (
		username VARCHAR(50) PRIMARY KEY,
		password VARCHAR(50) NOT NULL,
		firstname VARCHAR(50) NOT NULL,
		lastname VARCHAR(50) NOT NULL,
		mobile INT NOT NULL,
		userid VARCHAR(50) NOT NULL,
		profilepic VARCHAR(50)
	);
`
}
