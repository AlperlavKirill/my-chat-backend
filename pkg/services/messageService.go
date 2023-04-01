package services

import (
	"ChatProgramming/pkg/models"
	"ChatProgramming/pkg/session"
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageService struct {
	db    *mongo.Client
	mutex sync.Mutex
}

const (
	dbName         = "ChatBackend"
	collectionName = "Messages"
)

func NewMessageService(db *mongo.Client) *MessageService {
	return &MessageService{
		db:    db,
		mutex: sync.Mutex{},
	}
}

func (m *MessageService) Create(content string, author string) (string, error) {
	collection := m.db.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := &models.Message{
		Id:           session.RandomString(20),
		Content:      content,
		Author:       author,
		CreationDate: time.Now(),
	}
	_, err := collection.InsertOne(ctx, message)

	return message.Id, err
}

func (m *MessageService) GetById(id string) (*models.Message, error) {
	collection := m.db.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "id", Value: id}}
	opts := options.FindOne()
	var result *models.Message
	err := collection.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MessageService) GetByAuthor(author string) ([]*models.Message, error) {
	collection := m.db.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "author", Value: author}}
	opts := options.Find()
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var results []*models.Message
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
