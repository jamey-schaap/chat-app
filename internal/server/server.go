package server

import (
	"chat-app/internal/config"
	"chat-app/internal/database"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type server struct {
	port int
	db   *sql.DB
}

func NewServer() *http.Server {
	newServer := &server{
		port: config.GetConfig().Server.Port,
		db:   database.New(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
