package main

import (
	"ChatProgramming/config"
	"ChatProgramming/pkg/handlers"
	"ChatProgramming/pkg/services"
	"ChatProgramming/pkg/socket"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))
	r.POST("/message/create", messageHandler.CreateMessage)
	r.GET("/message/findById/:id", messageHandler.GetMessageById)
	r.GET("/message/findByUsername/:username", messageHandler.GetMessagesByAuthor)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.GET("/user/:username", userHandler.GetUserInfo)

	r.GET("/chat-socket", func(c *gin.Context) {
		socket.WsHandler(wsClient, c.Writer, c.Request)
	})

	runServer(r)
}

func runServer(r *gin.Engine) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	port := os.Getenv("PORT")

	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
