package main

import (
	"ChatProgramming/pkg/controllers/test"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
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
