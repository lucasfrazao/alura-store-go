package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connect := "user=postgres dbname=alura_store password= host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connect)

	if err != nil {
		panic(err.Error())
	}

	return db
}
