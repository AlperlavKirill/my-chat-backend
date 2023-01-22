package models

import "time"

type User struct {
	Username         string    `json:"username" bson:"username"`
	FirstName        string    `json:"firstname,omitempty" bson:"first_name"`
	LastName         string    `json:"lastname,omitempty" bson:"last_name"`
	PasswordHash     string    `json:"password_hash" bson:"password_hash"`
	PasswordSalt     string    `json:"password_salt" bson:"password_salt"`
	RegistrationDate time.Time `json:"registration_date" bson:"registration_date"`
}

type UserPublicInfo struct {
	Username         string
	FirstName        string
	LastName         string
	RegistrationDate time.Time
}

type UserRepository interface {
	Register(username, firstName, lastName, password string) (string, error)
	Login(username, password string) (string, error)
	GetInfo(username string) (*UserPublicInfo, error)
}
