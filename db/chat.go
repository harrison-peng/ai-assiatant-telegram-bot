package db

import (
	"context"
	"github.com/hichyen1207/ai-assiatant-telegram-bot/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Chat is a collection of chats.
type Chat struct {
	collection *mongo.Collection
}

// ChatCollection is a collection of chats.
func (db *MongoDB) ChatCollection() *Chat {
	return &Chat{
		collection: db.client.Database(database).Collection(collectionChat),
	}
}

// Insert inserts a chat into the database.
func (c *Chat) Insert(chatId int64, user models.User) error {
	ctx := context.Background()
	now := time.Now()

	if _, err := c.collection.InsertOne(ctx, bson.M{
		"chat_id":    chatId,
		"user":       user,
		"mode":       models.ModeChat,
		"language":   models.LanguageEnglish,
		"created_at": now,
		"updated_at": now,
	}); err != nil {
		return err
	}

	return nil
}

// SetChatGPTToken sets the chat GPT token.
func (c *Chat) SetChatGPTToken(chatId int64, chatGPTToken string) error {
	ctx := context.Background()

	filter := bson.M{"chat_id": chatId}
	update := bson.M{
		"$set": bson.M{
			"chat_gpt_token": chatGPTToken,
			"mode":           models.ModeChat,
			"updated_at":     time.Now(),
		},
	}

	if _, err := c.collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

// GetByChatId gets a message by chat id.
func (c *Chat) GetByChatId(chatId int64) (*models.Chat, error) {
	ctx := context.Background()

	filter := bson.M{"chat_id": chatId}
	var chat models.Chat
	if err := c.collection.FindOne(ctx, filter).Decode(&chat); err != nil {
		return nil, err
	}

	return &chat, nil
}

// UpdateMode updates the mode of a chat.
func (c *Chat) UpdateMode(chatId int64, mode models.Mode) error {
	ctx := context.Background()

	filter := bson.M{"chat_id": chatId}
	update := bson.M{
		"$set": bson.M{
			"mode":       mode,
			"updated_at": time.Now(),
		},
	}

	if _, err := c.collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

// UpdateLanguage updates the language of a chat.
func (c *Chat) UpdateLanguage(chatId int64, language models.Language) error {
	ctx := context.Background()

	filter := bson.M{"chat_id": chatId}
	update := bson.M{
		"$set": bson.M{
			"language":   language,
			"mode":       models.ModeChat,
			"updated_at": time.Now(),
		},
	}

	if _, err := c.collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}
