package server

import (
	"chat-app/internal/database"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type server struct {
	port int
	db   *sql.DB
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	fmt.Println(port)
	newServer := &server{
		port: port,
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
