package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type chatMessage struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	UserId  string `json:"userId"`
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
			// throw a specific error, catch it with middleware and return generic error
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
	}

	c.JSON(http.StatusOK, chat)
}

func (s *server) postChatHandler(c *gin.Context) {
	var newChatMessage chatMessage
	if err := c.BindJSON(&newChatMessage); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := s.db.Exec("INSERT INTO chat_messages VALUES (?, ?, ?)", newChatMessage.ID, newChatMessage.Message, newChatMessage.UserId)
	if err != nil {
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
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, newChatMessage)
}

func (s *server) RegisterRoutes() http.Handler {
	router := gin.Default()

	router.GET("/chats", s.getChatsHandler)
	router.GET("/chats/:id", s.getChatByIdHandler)
	router.POST("/chats", s.postChatHandler)
	router.PUT("/chats", s.updateChatHandler)

	return router
}
