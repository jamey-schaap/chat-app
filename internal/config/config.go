package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	Port int
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

	cfg = &Config{
		MySQL: MySQLConfig{
			User:   "root",
			Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
			DBName: os.Getenv("MYSQL_DATABASE"),
			Addr:   fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		},
		Server: ServerConfig{
			Port: port,
		},
	}

	return cfg
}
