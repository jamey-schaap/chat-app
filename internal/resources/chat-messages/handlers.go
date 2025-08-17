package chat_messages

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Controller struct {
	repository *Repository
}

func NewController(db *sql.DB) *Controller {
	return &Controller{
		repository: NewRepository(db),
	}
}

func (c *Controller) GetChatsHandler(w http.ResponseWriter, r *http.Request) {
	chatsMessages, err := c.repository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(chatsMessages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) GetChatByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := uuid.Parse(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatMessage, err := c.repository.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(chatMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) PostChatHandler(w http.ResponseWriter, r *http.Request) {
	var creationRequest CreateChatMessageRequest

	err := json.NewDecoder(r.Body).Decode(&creationRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tmpUserId, _ := uuid.Parse("aa48082a-5d5a-4147-9de3-2d994b6f790d") // TODO: Remove later
	newChatMessage := &ChatMessage{
		ID:        uuid.New(),
		Message:   creationRequest.Message,
		CreatedAt: time.Now().UTC(),
		UserId:    tmpUserId,
	}

	chatMessage, err := c.repository.Create(newChatMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(chatMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: mask all 500s later
		return
	}
}

func (c *Controller) UpdateChatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := uuid.Parse(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var chatMessage ChatMessage
	err = json.NewDecoder(r.Body).Decode(&chatMessage)
	if err != nil || chatMessage.ID != id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(updateRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
