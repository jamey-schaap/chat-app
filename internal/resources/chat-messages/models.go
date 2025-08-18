package chat_messages

import (
	"time"

	"github.com/google/uuid"
)

type ChatMessage struct {
	ID        uuid.UUID `json:"id" sql:"type:uuid"`
	Message   string    `json:"message"`
	UserId    uuid.UUID `json:"userId" db:"user_id" sql:"type:uuid"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" db:"updated_at"`
}

type CreateChatMessageRequest struct {
	Message string `json:"message"`
}

type PatchChatMessageRequest struct {
	Message string `json:"message"`
}
