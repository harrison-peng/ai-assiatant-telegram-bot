package db

import (
	"context"
	"fmt"
	"github.com/hichyen1207/ai-assiatant-telegram-bot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database = "telegramAIChat"

	collectionChat    = "chat"
	collectionMessage = "message"
)

// MongoDB is a MongoDB database.
type MongoDB struct {
	client *mongo.Client
}

// NewMongoDB creates a new MongoDB instance.
func NewMongoDB() (*MongoDB, error) {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority", config.MongoDbUsername, config.MongoDbPassword, config.MongoDbHost)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoDB{client: client}, nil
}
