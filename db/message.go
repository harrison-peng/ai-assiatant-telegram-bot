package db

import (
	"context"
	"github.com/hichyen1207/ai-assiatant-telegram-bot/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Message is a collection of messages.
type Message struct {
	collection *mongo.Collection
}

// MessageCollection is a collection of messages.
func (db *MongoDB) MessageCollection() *Message {
	return &Message{
		collection: db.client.Database(database).Collection(collectionMessage),
	}
}

// Insert inserts a message into the database.
func (c *Message) Insert(messageId int64, chatId int64, from int64, text string, timestamp int) error {
	ctx := context.Background()
	now := time.Now()

	if _, err := c.collection.InsertOne(ctx, bson.M{
		"message_id": messageId,
		"chat_id":    chatId,
		"from":       from,
		"text":       text,
		"timestamp":  timestamp,
		"archived":   false,
		"created_at": now,
		"updated_at": now,
	}); err != nil {
		return err
	}

	return nil
}

// GetLastMessages gets the last messages by chat id.
func (c *Message) GetLastMessages(chatId int64, count int) ([]models.Message, error) {
	ctx := context.Background()

	filter := bson.M{
		"chat_id":  chatId,
		"archived": false,
	}
	opts := options.Find().SetSort(bson.M{"created_at": -1}).SetLimit(int64(count))

	cursor, err := c.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

// Archive archives all messages by chat id.
func (c *Message) Archive(chatId int64) error {
	ctx := context.Background()

	filter := bson.M{
		"chat_id": chatId,
		"archived": bson.M{
			"$ne": true,
		},
	}
	update := bson.M{
		"$set": bson.M{
			"archived":   true,
			"updated_at": time.Now(),
		},
	}

	if _, err := c.collection.UpdateMany(ctx, filter, update); err != nil {
		return err
	}

	return nil
}
