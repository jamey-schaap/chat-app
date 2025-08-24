package server

import (
	"chat-app/internal/config"
	chatMessages "chat-app/internal/resources/chat-messages"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func RegisterRoutes(router *mux.Router, config *config.Config, db *sql.DB, logger *zap.Logger) http.Handler {
	chatsController := chatMessages.NewController(db, logger)

	router.HandleFunc("/chats", chatsController.GetChatsHandler).Methods("GET")
	router.HandleFunc("/chats/{id}", chatsController.GetChatByIdHandler).Methods("GET")
	router.HandleFunc("/chats", chatsController.PostChatHandler).Methods("POST")
	router.HandleFunc("/chats/{id}", chatsController.PatchChatHandler).Methods("PATCH")
	router.HandleFunc("/echo", chatsController.HandleWebSocketConnections)

	return router
}
