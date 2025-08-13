package config

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	MySQL *mysql.Config
}

type MySQLConfig struct {
	User   string
	Passwd string
	DBName string
	Addr   string
	Net    string
}

var (
	cfg *Config
)

func LoadConfig() *Config {
	if cfg != nil {
		return cfg
	}

	cfg = &Config{
		MySQL: mysql.NewConfig(),
	}

	cfg.MySQL.User = os.Getenv("MYSQL_USER")
	cfg.MySQL.Passwd = os.Getenv("MYSQL_PASSWD")
	cfg.MySQL.DBName = os.Getenv("MYSQL_DATABASE")
	cfg.MySQL.Addr = fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"))
	cfg.MySQL.Net = "tcp"

	return cfg
}
