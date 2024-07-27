package event

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type event struct {
	bot *tgbotapi.BotAPI
}

func NewEvent(bot *tgbotapi.BotAPI) *event {
	return &event{
		bot: bot,
	}
}

func (e *event) HandleEvent(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		e.sendWelcomeMessage(update)

		// handle when user send message
		if update.Message != nil {
			if update.Message.Text != "" {
				e.handleMessage(update)
			}
		}
		log.Printf("Got message %s from %s\n", update.Message.Text, update.Message.From)

	}
}
