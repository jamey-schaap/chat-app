package server

import (
	"chat-app/internal/config"
	"chat-app/internal/server/middleware"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewServer(config *config.Config, db *sql.DB, logger *zap.Logger) http.Handler {
	router := mux.NewRouter()
	RegisterRoutes(router, config, db, logger)

	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.LoggingMiddleware(logger))
	return router
}
