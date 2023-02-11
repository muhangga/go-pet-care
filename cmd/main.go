package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/muhangga/config"
	"github.com/muhangga/config/db"
	"github.com/muhangga/internal/app"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load config DB
	db, err := db.InitDB(true)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	config := config.NewConfig()
	server := app.InitServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		server.Run()
		wg.Done()
	}()

	wg.Wait()
}
