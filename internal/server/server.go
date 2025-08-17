package server

import (
	"chat-app/internal/config"
	"database/sql"
	"net/http"
)

func NewServer(config *config.Config, db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	registerRoutes(mux, config, db)

	var handler http.Handler = mux
	handler = CorsMiddleware(handler)
	return handler
}
