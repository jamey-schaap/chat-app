package models

import "github.com/google/uuid"

type CreateChatMessageRequest struct {
	Message string `json:"message"`
}

type UpdateChatMessageRequest struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}
