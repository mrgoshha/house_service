package postrges

import (
	"fmt"
	"os"
)

const (
	dbName   = "POSTGRES_DB"
	user     = "POSTGRES_USER"
	password = "POSTGRES_PASSWORD"
	ports    = "POSTGRES_PORTS"
	host     = "POSTGRES_HOST"
)

type Config struct {
	DbName   string
	User     string
	Password string
	Host     string
	Port     string
}

func NewConfig() (*Config, error) {
	dbName := os.Getenv(dbName)
	if len(dbName) == 0 {
		return nil, fmt.Errorf("config error")
	}
	user := os.Getenv(user)
	if len(user) == 0 {
		return nil, fmt.Errorf("config error")
	}
	password := os.Getenv(password)
	if len(password) == 0 {
		return nil, fmt.Errorf("config error")
	}
	ports := os.Getenv(ports)
	if len(ports) == 0 {
		return nil, fmt.Errorf("config error")
	}
	host := os.Getenv(host)
	if len(host) == 0 {
		return nil, fmt.Errorf("config error")
	}

	return &Config{
		DbName:   dbName,
		User:     user,
		Password: password,
		Host:     host,
		Port:     ports,
	}, nil
}
