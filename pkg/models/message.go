package models

import "time"

type Message struct {
	Id           int       `json:"id"`
	Content      string    `json:"content"`
	Author       string    `json:"author"` //username
	CreationDate time.Time `json:"creation_date"`
}

type MessageRepository interface {
	Create(content string, author string) (int, error)
	GetById(id int) (*Message, error)
	GetByAuthor(Author string) ([]*Message, error)
}
