package telegram

import (
	"github.com/hichyen1207/ai-assiatant-telegram-bot/db"
	"github.com/hichyen1207/ai-assiatant-telegram-bot/models"
)

// HandleCommand handles the command.
func (tg *Telegram) HandleCommand() (string, error) {
	message := tg.update.Message

	switch message.Command() {
	case "help":
		msg := `This bot is powered by ChatGPT. You can use the following commands:
/help - show this message
/settoken - set chat gpt token
/newtopic - start a new topic
/setlanguage - set language
`
		return msg, nil
	case "settoken":
		mongodb, err := db.NewMongoDB()
		if err != nil {
			return "", err
		}

		if err := mongodb.ChatCollection().UpdateMode(message.Chat.ID, models.ModeSetToken); err != nil {
			return "", err
		}

		return "Please send me your chat gpt token", nil
	case "newtopic":
		mongodb, err := db.NewMongoDB()
		if err != nil {
			return "", err
		}

		// archive messages
		if err := mongodb.MessageCollection().Archive(message.Chat.ID); err != nil {
			return "", err
		}

		return "You can start a new topic", nil
	case "setlanguage":
		mongodb, err := db.NewMongoDB()
		if err != nil {
			return "", err
		}

		chat, err := mongodb.ChatCollection().GetByChatId(message.Chat.ID)
		if err != nil {
			return "", err
		}

		if err := mongodb.ChatCollection().UpdateMode(message.Chat.ID, models.ModeSetLanguage); err != nil {
			return "", err
		}

		var msg string
		switch chat.Language {
		case models.LanguageEnglish:
			msg = `Your language is English. Please send the language you want to change to:
en - English
zh - Chinese`
		case models.LanguageMandarin:
			msg = `你的語言是中文。請發送你想要更改的語言：
en - 英文
zh - 中文`
		}

		return msg, nil
	default:
		return "I don't know that command", nil
	}
}
