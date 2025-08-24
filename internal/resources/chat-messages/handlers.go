package chat_messages

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Controller struct {
	chatMessageRepository *ChatMessageRepository
	logger                *zap.Logger
}

func NewController(db *sql.DB, logger *zap.Logger) *Controller {
	return &Controller{
		chatMessageRepository: NewChatMessageRepository(db),
		logger:                logger,
	}
}

func (c *Controller) GetChatsHandler(w http.ResponseWriter, _ *http.Request) {
	chatsMessages, err := c.chatMessageRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	chatMessage, err := c.chatMessageRepository.GetById(id)
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

	chatMessage, err := c.chatMessageRepository.Create(newChatMessage)
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

func (c *Controller) PatchChatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := uuid.Parse(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var patch PatchChatMessageRequest
	err = json.NewDecoder(r.Body).Decode(&patch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.chatMessageRepository.Patch(id, &patch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// todo implement
		return true
	},
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func (c *Controller) HandleWebSocketConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		c.logger.Error(err.Error(), zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			c.logger.Error(err.Error(), zap.Error(err))
			return
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			c.logger.Error(err.Error(), zap.Error(err))
			break
		}

		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			c.logger.Error(err.Error(), zap.Error(err))
			break
		}
	}
}
