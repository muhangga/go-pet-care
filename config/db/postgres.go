package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func InitDB(setLimits bool) (*sql.DB, error) {

	dsn := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
	}

	if setLimits {
		fmt.Println("Setting database limits")
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Error().Msgf("cant ping database %s", err)
	}

	fmt.Println("Database connected")

	return db, nil
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Error().Msgf("cant close database %s", err)
	}
}
