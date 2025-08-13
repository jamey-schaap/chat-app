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

func (s *server) getChats(c *gin.Context) {
	results, err := s.db.Query("SELECT * FROM chats")
	if err != nil {
		log.Fatal(err)
	}

	chats := make([]chatMessage, 0)
	for results.Next() {
		var chat chatMessage
		err = results.Scan(&chat.ID, &chat.Message, &chat.UserId)
		if err != nil {
			c.Status(http.StatusNotFound)
			// throw specific error, catch it with middleware and return generic error
		}

		chats = append(chats, chat)
	}

	c.JSON(http.StatusOK, chats)
}

func (s *server) getChatById(c *gin.Context) {
	id := c.Param("id")

	result := s.db.QueryRow("SELECT * FROM chats WHERE id = ?", id)
	var chat chatMessage

	err := result.Scan(&chat.ID, &chat.Message, &chat.UserId)
	if err != nil {
		c.Status(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, chat)
}

func (s *server) postChat(c *gin.Context) {
	var newChatMessage chatMessage
	if err := c.BindJSON(&newChatMessage); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := s.db.Exec("INSERT INTO chats VALUES (?, ?, ?)", newChatMessage.ID, newChatMessage.Message, newChatMessage.UserId)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, newChatMessage)
}

func (s *server) updateChat(c *gin.Context) {
	var newChatMessage chatMessage
	if err := c.BindJSON(&newChatMessage); err != nil {
		return
	}

	_, err := s.db.Exec("UPDATE users SET message = ? WHERE id = ?", newChatMessage.Message, newChatMessage.UserId)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, newChatMessage)
}

func (s *server) RegisterRoutes() http.Handler {
	router := gin.Default()

	router.GET("/chats", s.getChats)
	router.GET("/chats/:id", s.getChatById)
	router.POST("/chats", s.postChat)
	router.PUT("/chats", s.updateChat)

	return router
}
