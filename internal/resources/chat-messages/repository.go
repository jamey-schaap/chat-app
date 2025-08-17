package chat_messages

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll() ([]ChatMessage, error) {
	results, err := r.db.Query("SELECT * FROM chat_messages")
	if err != nil {
		return nil, err
	}

	messages := make([]ChatMessage, 0)
	for results.Next() {
		var chat ChatMessage
		err = results.Scan(&chat.ID, &chat.Message, &chat.UserId, &chat.CreatedAt, &chat.UpdatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, chat)
	}

	return messages, nil
}

func (r *Repository) GetById(id uuid.UUID) (*ChatMessage, error) {
	result := r.db.QueryRow("SELECT * FROM chat_messages WHERE id = uuid_to_bin(?)", id)

	var chatMessage ChatMessage
	err := result.Scan(&chatMessage.ID, &chatMessage.Message, &chatMessage.UserId, &chatMessage.CreatedAt, &chatMessage.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &chatMessage, nil
}

func (r *Repository) Create(chatMessage *ChatMessage) (*ChatMessage, error) {
	_, err := r.db.Exec("INSERT INTO chat_messages VALUES (UUID_TO_BIN(?), ?, uuid_to_bin(?))", chatMessage.ID, chatMessage.Message, chatMessage.UserId)
	if err != nil {
		return nil, err
	}

	return chatMessage, err
}

func (r *Repository) Update(chatMessage *ChatMessage) (*ChatMessage, error) {
	_, err := r.db.Exec("UPDATE chat_messages SET message = ? WHERE id = ?", updateRequest.Message, updateRequest.ID)
	if err != nil {
		return nil, err
	}
	return chatMessage, err
}
