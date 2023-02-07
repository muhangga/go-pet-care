package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/muhangga/config"
	"github.com/muhangga/internal/app"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.NewConfig()
	config.Database()

	server := app.InitServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		server.Run()
		wg.Done()
	}()

	wg.Wait()
}
