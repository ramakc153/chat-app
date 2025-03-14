package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	conn_string := "host=localhost port=5432 user=postgres password=admin dbname=chat-app sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", conn_string)
	if err != nil {
		panic(err)
	}
}
