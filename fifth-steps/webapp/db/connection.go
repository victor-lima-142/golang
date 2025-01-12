package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectDB() (*sql.DB, error) {
	conectionUrl := "user=postgres dbname=goproducts password=CYvr9tNwEbalWAZPsMiC host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conectionUrl)
	if err != nil {
		fmt.Println("Error in db connection:", err.Error())
		panic(err)
	}
	return db, err
}

func RunSQL(db *sql.DB, query string) sql.Result {
	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	return result
}
