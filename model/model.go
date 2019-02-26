package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func DBConnect() (db *sqlx.DB) {
	db, err := sqlx.Open("mysql", "gestbook:password@tcp(localhost:3306)/gestbook")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
