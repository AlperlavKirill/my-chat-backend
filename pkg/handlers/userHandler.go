package handlers

import (
	"ChatProgramming/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserRepo models.UserRepository
}

func (u *UserHandler) Register(c *gin.Context) {
	username := c.PostForm("username")
	firstName := c.DefaultPostForm("firstname", "")
	lastName := c.DefaultPostForm("lastname", "")
	password := c.PostForm("password")

	_, err := u.UserRepo.Register(username, firstName, lastName, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, username)
}

func (u *UserHandler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	_, err := u.UserRepo.Login(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, username)
}

func (u *UserHandler) GetUserInfo(c *gin.Context) {
	username := c.Param("username")

	userInfo, err := u.UserRepo.GetInfo(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, userInfo)
}
