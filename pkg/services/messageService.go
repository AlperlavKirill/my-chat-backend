package services

import (
	"ChatProgramming/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type MessageService struct {
	db    *mongo.Client
	mutex sync.Mutex
}

func (m *MessageService) Create(content string, author string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageService) GetById(id int) (*models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageService) GetByAuthor(Author string) ([]*models.Message, error) {
	//TODO implement me
	panic("implement me")
}
