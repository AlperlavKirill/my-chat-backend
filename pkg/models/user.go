package models

import "time"

type User struct {
	Username         string
	Password         string
	RegistrationDate time.Time
}

type UserRepository interface {
	Register(username string, password string) (string, error)
	GetInfo(username string) *User
}
