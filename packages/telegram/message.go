package telegram

import (
	"github.com/hichyen1207/ai-assiatant-telegram-bot/db"
	"github.com/hichyen1207/ai-assiatant-telegram-bot/models"
	"github.com/hichyen1207/ai-assiatant-telegram-bot/packages/openai"
	goopenai "github.com/sashabaranov/go-openai"
	"strings"
	"time"
)

// HandleMessage handles a message.
func (tg *Telegram) HandleMessage() (string, error) {
	mongodb, err := db.NewMongoDB()
	if err != nil {
		return "", err
	}

	message := tg.update.Message
	chat, err := mongodb.ChatCollection().GetByChatId(message.Chat.ID)
	if err != nil {
		if !db.IsNotFoundError(err) {
			return "", err
		}

		chat = &models.Chat{
			ChatId: message.Chat.ID,
			User: models.User{
				UserId:    message.From.ID,
				UserName:  message.From.UserName,
				FirstName: message.From.FirstName,
			},
		}
		if err := mongodb.ChatCollection().Insert(chat.ChatId, chat.User); err != nil {
			return "", err
		}
	}

	switch chat.Mode {
	case models.ModeChat:
		if chat.ChatGPTToken == nil {
			var msg string
			switch chat.Language {
			case models.LanguageEnglish:
				msg = "Please add chat gpt token to start conversation"
			case models.LanguageMandarin:
				msg = "請設定 Chat GPT token 來開始對話"
			}
			return msg, nil
		}

		msgs, err := mongodb.MessageCollection().GetLastMessages(chat.ChatId, 20)
		if err != nil {
			return "", err
		}

		userRole := "system"
		assistantRole := "assistant"
		var messages []goopenai.ChatCompletionMessage
		switch chat.Language {
		case models.LanguageEnglish:
			messages = append(messages, goopenai.ChatCompletionMessage{
				Role:    "system",
				Content: "You are a helpful assistant.",
			})
		case models.LanguageMandarin:
			messages = append(messages, goopenai.ChatCompletionMessage{
				Role:    "system",
				Content: "你是一個的智能助理，可以回答使用者的任何問題。",
			})
		}

		if len(msgs) > 0 {
			if time.Now().After(time.Unix(int64(msgs[0].Timestamp), 0).Add(time.Hour)) {
				// archive messages
				if err := mongodb.MessageCollection().Archive(message.Chat.ID); err != nil {
					return "", err
				}

				messages = append(messages, goopenai.ChatCompletionMessage{
					Role:    userRole,
					Content: message.Text,
				})
			} else {
				for i := len(msgs) - 1; i >= 0; i-- {
					msg := msgs[i]

					sender := userRole
					if msg.From == 0 {
						sender = assistantRole
					}

					messages = append(messages, goopenai.ChatCompletionMessage{
						Role:    sender,
						Content: msg.Text,
					})
				}

				messages = append(messages, goopenai.ChatCompletionMessage{
					Role:    userRole,
					Content: message.Text,
				})
			}
		} else {
			messages = append(messages, goopenai.ChatCompletionMessage{
				Role:    userRole,
				Content: message.Text,
			})
		}

		if err := mongodb.MessageCollection().Insert(int64(message.MessageID), chat.ChatId, chat.User.UserId, message.Text, message.Date); err != nil {
			return "", err
		}

		client := openai.NewOpenAI(*chat.ChatGPTToken)
		response, err := client.Chat(messages)
		if err != nil {
			return "", err
		}

		response = strings.Trim(response, "\n ")

		if response == "" {
			if err := mongodb.MessageCollection().Archive(chat.ChatId); err != nil {
				return "", err
			}

			return "Thanks", nil
		}

		if err := mongodb.MessageCollection().Insert(int64(message.MessageID+1), chat.ChatId, 0, response, int(time.Now().Unix())); err != nil {
			return "", err
		}

		return response, nil
	case models.ModeSetToken:
		if err := mongodb.ChatCollection().SetChatGPTToken(chat.ChatId, message.Text); err != nil {
			return "", err
		}

		return "Chat gpt token has been set. Now you can start conversation with AI assistant.", nil
	case models.ModeSetLanguage:
		language := models.Language(message.Text)
		if !language.Validate() {
			if err := mongodb.ChatCollection().UpdateMode(message.Chat.ID, models.ModeChat); err != nil {
				return "", err
			}

			return "Invalid language", nil
		}

		if err := mongodb.ChatCollection().UpdateLanguage(message.Chat.ID, language); err != nil {
			return "", err
		}

		if err := mongodb.MessageCollection().Archive(chat.ChatId); err != nil {
			return "", err
		}

		switch language {
		case models.LanguageEnglish:
			return "Language has been set to English", nil
		case models.LanguageMandarin:
			return "語言已設定為中文", nil
		}
	}

	return "", nil
}
