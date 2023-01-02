package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SimpleHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, Guest")
}

func ParameterHello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello, %s", name)
}

func QueryHello(c *gin.Context) {
	firstName := c.DefaultQuery("firstName", "Guest")
	lastName := c.Query("lastName")

	c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
}

func PostBodyHello(c *gin.Context) {
	name := c.PostForm("name")
	c.String(http.StatusOK, "Hello %s", name)
}
