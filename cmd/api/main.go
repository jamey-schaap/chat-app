package main

import (
	"chat-app/internal/config"
	"chat-app/internal/database"
	"chat-app/internal/server"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

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
		Addr:         net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.Port)),
		Handler:      srv,
		ReadTimeout:  cfg.Server.ReadTimeout * time.Second,
		WriteTimeout: cfg.Server.WriteTimeout * time.Second,
		IdleTimeout:  cfg.Server.IdleTimeout * time.Second,
	}

	log.Printf("Listening on %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error listening and serving: %s\n", err)
	}
}
