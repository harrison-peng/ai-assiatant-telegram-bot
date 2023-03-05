package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hichyen1207/ai-assiatant-telegram-bot/config"
	"github.com/hichyen1207/ai-assiatant-telegram-bot/packages/telegram"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.TgBotToken)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		tg := telegram.NewTelegram(update)
		// handle message
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			if update.Message.IsCommand() {
				responseMsg, err := tg.HandleCommand()
				if err != nil {
					msg.Text = "Sorry, something went wrong. Please try again later."
				}

				msg.Text = responseMsg
			} else {
				respondMsg, err := tg.HandleMessage()
				if err != nil {
					msg.Text = "Sorry, something went wrong. Please try again later."
				}

				msg.Text = respondMsg
			}

			bot.Send(msg)
		}

	}
}
