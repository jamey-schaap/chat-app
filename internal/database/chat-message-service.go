package database

import (
	"chat-app/internal/models"

	"github.com/google/uuid"
)

type ChatMessageService interface {
	Get(id uuid.UUID) models.ChatMessage
}
