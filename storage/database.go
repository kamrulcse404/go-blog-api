package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {

	var err error
	connectionString := "host=localhost port=5432 user=bloguser password=blogpass dbname=blogdb sslmode=disable"

	DB, err = sql.Open(
		"postgres",
		connectionString,
	)

	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}
