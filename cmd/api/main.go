package main

import (
	"fmt"
	"net/http"
	"os"
	"slices"

	"chat-app/internal/database"

	"github.com/gin-gonic/gin"
)

var chats = []chatMessage{
	{ID: "1", Message: "Hello World", UserId: "123"},
	{ID: "2", Message: "Hello World", UserId: "321"},
	{ID: "3", Message: "Hello World", UserId: "213"},
}

type chatMessage struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	UserId  string `json:"userId"`
}

func getChats(c *gin.Context) {
	c.JSON(http.StatusOK, chats)
}

func getChatById(c *gin.Context) {
	id := c.Param("id")
	index := slices.IndexFunc(chats, func(chat chatMessage) bool {
		return chat.ID == id
	})

	if index == -1 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, chats[index])
}

func postChat(c *gin.Context) {
	var newChatMessage chatMessage
	if err := c.BindJSON(&newChatMessage); err != nil {
		return
	}

	chats = append(chats, newChatMessage)
	c.JSON(http.StatusOK, newChatMessage)
}

func updateChat(c *gin.Context) {
	var newChatMessage chatMessage
	if err := c.BindJSON(&newChatMessage); err != nil {
		return
	}

	index := slices.IndexFunc(chats, func(chat chatMessage) bool {
		return chat.ID == newChatMessage.ID
	})

	if index == -1 {
		c.Status(http.StatusNotFound)
		return
	}

	chats[index] = newChatMessage
	c.JSON(http.StatusOK, newChatMessage)
}

func main() {
	db := database.ConnectToMySql()

	router := gin.Default()
	router.GET("/chats", getChats)
	router.GET("/chats/:id", getChatById)
	router.POST("/chats", postChat)
	router.PUT("/chats", updateChat)

	addr := fmt.Sprintf("localhost:%s", os.Getenv("PORT"))
	err := router.Run(addr)
	if err != nil {
		panic(err)
	}
}
