package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var PSG *sql.DB

func ConnectToDB(conn string) error {
	var err error
	PSG, err = sql.Open("postgres", conn)
	if err != nil {
		logrus.Errof("Unable to connect to postgres db %v", err)
		return err
	}
	return nil
}
