package main

import (
	"ChatProgramming/config"
	"ChatProgramming/pkg/handlers"
	"ChatProgramming/pkg/handlers/test"
	"ChatProgramming/pkg/services"
	"ChatProgramming/pkg/socket"
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

	userService := services.NewUserService(config.PostgresDB())
	userHandler := &handlers.UserHandler{
		UserRepo: userService,
	}

	wsClient := &socket.Client{
		MessageRepo: messageService,
	}

	r := gin.Default()
	r.POST("/message/create", messageHandler.CreateMessage)
	r.GET("/message/findById/:id", messageHandler.GetMessageById)
	r.GET("/message/findByUsername/:username", messageHandler.GetMessagesByAuthor)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.GET("/user/:username", userHandler.GetUserInfo)

	r.GET("/chat-socket", func(c *gin.Context) {
		socket.WsHandler(wsClient, c.Writer, c.Request)
	})

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
