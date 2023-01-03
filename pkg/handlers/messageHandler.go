package handlers

import (
	"ChatProgramming/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageHandler struct {
	MessageRepo models.MessageRepository
}

func (m *MessageHandler) CreateMessage(c *gin.Context) {
	content := c.PostForm("content")
	author := c.PostForm("author")

	_, err := m.MessageRepo.Create(content, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
}

func (m *MessageHandler) GetMessageById(c *gin.Context) {
	id := c.Param("id")

	message, err := m.MessageRepo.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, message)
}

func (m *MessageHandler) GetMessagesByAuthor(c *gin.Context) {
	author := c.Param("username")

	messages, err := m.MessageRepo.GetByAuthor(author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, messages)
}
