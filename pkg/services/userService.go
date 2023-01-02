package services

import (
	"ChatProgramming/pkg/models"
	"database/sql"
	"sync"
)

type UserService struct {
	db    *sql.DB
	mutex sync.Mutex
}

func (u *UserService) Register(username string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetInfo(username string) *models.User {
	//TODO implement me
	panic("implement me")
}
