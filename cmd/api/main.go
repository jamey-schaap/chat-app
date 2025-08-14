package main

import (
	"chat-app/internal/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") == "local" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}

	s := server.NewServer()
	// defer db.close()
	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %v", err))
	}
}
