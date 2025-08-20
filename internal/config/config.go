package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	MySQL  MySQLConfig
	Server ServerConfig
}

type MySQLConfig struct {
	User   string
	Passwd string
	DBName string
	Addr   string
}

type ServerConfig struct {
	Host         string
	Port         int
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var (
	cfg *Config
)

func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	idleTimeout, err := time.ParseDuration(os.Getenv("IDLE_TIMEOUT"))
	if err != nil {
		log.Fatal(err)
	}

	readTimeout, err := time.ParseDuration(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		log.Fatal(err)
	}

	writeTimeout, err := time.ParseDuration(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		log.Fatal(err)
	}

	cfg = &Config{
		MySQL: MySQLConfig{
			User:   "root",
			Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
			DBName: os.Getenv("MYSQL_DATABASE"),
			Addr:   fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		},
		Server: ServerConfig{
			Port:         port,
			IdleTimeout:  idleTimeout,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	}

	return cfg
}
