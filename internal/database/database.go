package database

import (
	"chat-app/internal/config"
	"database/sql"
	"fmt"
	"log"
)

var (
	db *sql.DB
)

func New() *sql.DB {
	if db != nil {
		return db
	}

	cfg := config.GetConfig().MySQL

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}
