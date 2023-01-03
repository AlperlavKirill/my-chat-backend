package main

import (
	"ChatProgramming/config"
	"ChatProgramming/pkg/handlers"
	"ChatProgramming/pkg/handlers/test"
	"ChatProgramming/pkg/services"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	messageService := services.NewMessageService(config.MongoDB())
	messageHandler := &handlers.MessageHandler{
		MessageRepo: messageService,
	}

	r := gin.Default()
	r.POST("/message/create", messageHandler.CreateMessage)
	r.GET("/message/findById/:id", messageHandler.GetMessageById)
	r.GET("/message/findByUsername/:username", messageHandler.GetMessagesByAuthor)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/simple", test.SimpleHello)
	r.GET("/parameter/:name", test.ParameterHello)
	r.GET("/query", test.QueryHello)
	r.POST("/body", test.PostBodyHello)

	err := r.Run(":8080")
	if err != nil {
		//do nothing yet
	}
}
