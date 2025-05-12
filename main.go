package main

import (
	"backend-challenge/api/routes"
	"backend-challenge/config"
	"log"
)

func main() {

	// Load config before anything else
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// API mode
	r := routes.SetupRouter()
	log.Println("Starting server on :8080")
	r.Run(":8080")
}
