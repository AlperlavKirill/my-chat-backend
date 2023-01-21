package models

import "time"

type Message struct {
	Id           string    `json:"id" bson:"id"`
	Content      string    `json:"content" bson:"content"`
	Author       string    `json:"author" bson:"author"` //username
	CreationDate time.Time `json:"creation_date" bson:"creation_date"`
}

type MessageRepository interface {
	Create(content, author string) (string, error)
	GetById(id string) (*Message, error)
	GetByAuthor(author string) ([]*Message, error)
}
