package main

import (
	"chat-app/internal/config"
	"chat-app/internal/database"
	"chat-app/internal/server"
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	if env := os.Getenv("APP_ENV"); env == "local" || env == "development" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
		logger = zap.Must(zap.NewDevelopment())
	}

	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	cfg := config.GetConfig()
	db := database.New(cfg.MySQL)
	srv := server.NewServer(cfg, db, logger)

	httpServer := &http.Server{
		Addr:         net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.Port)),
		Handler:      srv,
		ReadTimeout:  cfg.Server.ReadTimeout * time.Second,
		WriteTimeout: cfg.Server.WriteTimeout * time.Second,
		IdleTimeout:  cfg.Server.IdleTimeout * time.Second,
	}

	serverError := make(chan error, 1)
	go func() {
		log.Printf("Listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverError <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	select {
	case err := <-serverError:
		log.Printf("Server error: %v", err)
	case sig := <-stop:
		log.Printf("Received shutdown signal: %v", sig)
	}

	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown with error: %v", err)
		return
	}

	log.Println("Server exited gracefully")
}
