package server

import (
	"chat-app/internal/config"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(config *config.Config, db *sql.DB) http.Handler {
	router := mux.NewRouter()
	RegisterRoutes(router, config, db)

	router.Use(mux.CORSMethodMiddleware(router))
	return router
}
