package models

import "time"

type Message struct {
	Id           int
	Content      string
	Author       string //username
	CreationDate time.Time
}

type MessageRepository interface {
	GetById(id int) (*Message, error)
	GetByAuthor(Author string) ([]*Message, error)
}
