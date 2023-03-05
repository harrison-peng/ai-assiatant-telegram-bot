package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Telegram struct {
	update tgbotapi.Update
}

func NewTelegram(update tgbotapi.Update) *Telegram {
	return &Telegram{
		update: update,
	}
}
