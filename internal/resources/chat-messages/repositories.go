package chat_messages

import (
	"database/sql"

	"github.com/google/uuid"
)

type ChatMessageRepository struct {
	db *sql.DB
}

func NewChatMessageRepository(db *sql.DB) *ChatMessageRepository {
	return &ChatMessageRepository{
		db: db,
	}
}

func (r *ChatMessageRepository) GetAll() ([]ChatMessage, error) {
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

func (r *ChatMessageRepository) GetById(id uuid.UUID) (*ChatMessage, error) {
	result := r.db.QueryRow("SELECT * FROM chat_messages WHERE id = UUID_TO_BIN(?)", id)

	var chatMessage ChatMessage
	err := result.Scan(&chatMessage.ID, &chatMessage.Message, &chatMessage.UserId, &chatMessage.CreatedAt, &chatMessage.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &chatMessage, nil
}

func (r *ChatMessageRepository) Create(chatMessage *ChatMessage) (*ChatMessage, error) {
	_, err := r.db.Exec("INSERT INTO chat_messages (id, message, user_id, created_at) VALUES (UUID_TO_BIN(?), ?, UUID_TO_BIN(?), ?)", chatMessage.ID, chatMessage.Message, chatMessage.UserId, chatMessage.CreatedAt)
	if err != nil {
		return nil, err
	}

	return chatMessage, err
}

func (r *ChatMessageRepository) Patch(id uuid.UUID, patch *PatchChatMessageRequest) error {
	_, err := r.db.Exec("UPDATE chat_messages SET message = ? WHERE id = ?", patch.Message, id)
	if err != nil {
		return err
	}
	return err
}
