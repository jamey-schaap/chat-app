package server

import (
	chat_messages "chat-app/internal/resources/chat-messages"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(db *sql.DB) http.Handler {
	router := mux.NewRouter()

	chatsController := chat_messages.NewController(db)

	router.HandleFunc("/chats", chatsController.GetChatsHandler).Methods("GET")
	router.HandleFunc("/chats/{id}", chatsController.GetChatByIdHandler).Methods("GET")
	router.HandleFunc("/chats", chatsController.PostChatHandler).Methods("POST")
	router.HandleFunc("/chats/{id}", chatsController.UpdateChatHandler).Methods("PUT")

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
