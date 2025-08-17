package database

import (
	"chat-app/internal/config"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func New(cfg config.MySQLConfig) *sql.DB {
	if db != nil {
		return db
	}

	mySqlConfig := mysql.Config{
		User:      cfg.User,
		Passwd:    cfg.Passwd,
		DBName:    cfg.DBName,
		Addr:      cfg.Addr,
		Net:       "tcp",
		ParseTime: true,
	}

	var err error
	db, err = sql.Open("mysql", mySqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
