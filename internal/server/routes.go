package server

import (
	"chat-app/internal/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type chatMessage struct {
	ID        uuid.UUID `json:"id" sql:"type:uuid"`
	Message   string    `json:"message"`
	UserId    uuid.UUID `json:"userId"  sql:"type:uuid"`
	TimeStamp time.Time `json:"timestamp"`
}

func (s *server) getChatsHandler(c *gin.Context) {
	results, err := s.db.Query("SELECT * FROM chat_messages")
	if err != nil {
		log.Fatal(err)
	}

	chats := make([]chatMessage, 0)
	for results.Next() {
		var chat chatMessage
		err = results.Scan(&chat.ID, &chat.Message, &chat.UserId)
		if err != nil {
			c.Status(http.StatusNotFound)
			// throw a specific error, catch it with middleware and return generic error?
			return
		}

		chats = append(chats, chat)
	}

	c.JSON(http.StatusOK, chats)
}

func (s *server) getChatByIdHandler(c *gin.Context) {
	id := c.Param("id")

	result := s.db.QueryRow("SELECT * FROM chat_messages WHERE id = ?", id)
	var chat chatMessage

	err := result.Scan(&chat.ID, &chat.Message, &chat.UserId)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, chat)
}

func (s *server) postChatHandler(c *gin.Context) {
	var request models.CreateChatMessageRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	tmpUserId, _ := uuid.Parse("aa48082a-5d5a-4147-9de3-2d994b6f790d")
	newChatMessage := chatMessage{
		ID:        uuid.New(),
		Message:   request.Message,
		TimeStamp: time.Now().UTC(),
		UserId:    tmpUserId,
	}

	_, err := s.db.Exec("INSERT INTO chat_messages VALUES (UUID_TO_BIN(?), ?, uuid_to_bin(?))", newChatMessage.ID, newChatMessage.Message, newChatMessage.UserId)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, newChatMessage)
}

func (s *server) updateChatHandler(c *gin.Context) {
	var newChatMessage chatMessage
	if err := c.BindJSON(&newChatMessage); err != nil {
		return
	}

	_, err := s.db.Exec("UPDATE chat_messages SET message = ? WHERE id = ?", newChatMessage.Message, newChatMessage.UserId)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, newChatMessage)
}

func (s *server) RegisterRoutes() http.Handler {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/chats", s.getChatsHandler)
	router.GET("/chats/:id", s.getChatByIdHandler)
	router.POST("/chats", s.postChatHandler)
	router.PUT("/chats", s.updateChatHandler)

	return router
}
