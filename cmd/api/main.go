package main

import (
	"chat-app/internal/config"
	"chat-app/internal/database"
	"chat-app/internal/server"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") == "local" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}

	cfg := config.GetConfig()
	db := database.New(cfg.MySQL)
	srv := server.NewServer(cfg, db)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort("host", "port"),
		Handler: srv,
	}

	go func() {
		log.Printf("Listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			_, _ = fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
}
