package models

import "time"

type User struct {
	Username         string    `json:"username" bson:"username"`
	Password         string    `json:"password" bson:"password"`
	RegistrationDate time.Time `json:"registration_date" bson:"registration_date"`
}

type UserRepository interface {
	Register(username string, password string) (string, error)
	GetInfo(username string) *User
}
