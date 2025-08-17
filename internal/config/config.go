package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	MySQL  *mysql.Config
	Server ServerConfig
}

type MySQLConfig struct {
	User   string
	Passwd string
	DBName string
	Addr   string
	Net    string
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
		MySQL: mysql.NewConfig(),
		Server: ServerConfig{
			Port: port,
		},
	}

	cfg.MySQL.User = "root"
	cfg.MySQL.Passwd = os.Getenv("MYSQL_ROOT_PASSWORD")
	cfg.MySQL.DBName = os.Getenv("MYSQL_DATABASE")
	cfg.MySQL.Addr = fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"))
	cfg.MySQL.Net = "tcp"
	cfg.MySQL.ParseTime = true

	return cfg
}
