package server

import (
	"chat-app/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type chatMessage struct {
	ID        uuid.UUID `json:"id" sql:"type:uuid"`
	Message   string    `json:"message"`
	UserId    uuid.UUID `json:"userId"  sql:"type:uuid"`
	TimeStamp time.Time `json:"timestamp"`
}

func (s *server) getChatsHandler(w http.ResponseWriter, r *http.Request) {
	results, err := s.db.Query("SELECT * FROM chat_messages")
	if err != nil {
		log.Fatal(err)
	}

	messages := make([]chatMessage, 0)
	for results.Next() {
		var chat chatMessage
		err = results.Scan(&chat.ID, &chat.Message, &chat.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			// throw a specific error, catch it with middleware and return generic error?
			return
		}

		messages = append(messages, chat)
	}

	w.Header().Set("Content-Type", "applicaton/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) getChatByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := uuid.Parse(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := s.db.QueryRow("SELECT * FROM chat_messages WHERE id = uuid_to_bin(?)", id)

	var chatMessage chatMessage
	err = result.Scan(&chatMessage.ID, &chatMessage.Message, &chatMessage.UserId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(chatMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) postChatHandler(w http.ResponseWriter, r *http.Request) {
	var createChatMessageRequest models.CreateChatMessageRequest

	err := json.NewDecoder(r.Body).Decode(&createChatMessageRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tmpUserId, _ := uuid.Parse("aa48082a-5d5a-4147-9de3-2d994b6f790d") // TODO: Remove later
	newChatMessage := chatMessage{
		ID:        uuid.New(),
		Message:   createChatMessageRequest.Message,
		TimeStamp: time.Now().UTC(),
		UserId:    tmpUserId,
	}

	_, err = s.db.Exec("INSERT INTO chat_messages VALUES (UUID_TO_BIN(?), ?, uuid_to_bin(?))", newChatMessage.ID, newChatMessage.Message, newChatMessage.UserId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(newChatMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: mask all 500s later
		return
	}
}

func (s *server) updateChatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := uuid.Parse(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newChatMessage chatMessage
	err = json.NewDecoder(r.Body).Decode(&newChatMessage)
	if err != nil || newChatMessage.ID != id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = s.db.Exec("UPDATE chat_messages SET message = ? WHERE id = ?", newChatMessage.Message, newChatMessage.UserId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(newChatMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) RegisterRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/chats", s.getChatsHandler).Methods("GET")
	router.HandleFunc("/chats/{id}", s.getChatByIdHandler).Methods("GET")
	router.HandleFunc("/chats", s.postChatHandler).Methods("POST")
	router.HandleFunc("/chats/{id}", s.updateChatHandler).Methods("PUT")

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
