package config

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/muhangga/config/db"
)

type (
	config struct {
	}

	Config interface {
		ServiceName() string
		ServicePort() int
		ServiceEnvironment() string
		Database() *sql.DB
	}
)

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *sql.DB {
	db, err := db.InitDB(true)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func (c *config) ServicePort() int {
	v := os.Getenv("SERVICE_PORT")
	port, _ := strconv.Atoi(v)

	return port
}

func (c *config) ServiceEnvironment() string {
	return os.Getenv("ENV")
}
