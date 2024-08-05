package event

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (e *event) handleMessage(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID
	if _, err := e.bot.Send(msg); err != nil {
		panic(err)
	}
}

func (e *event) sendWelcomeMessage(update tgbotapi.Update, user string) {
	fmt.Println("user", user)
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.ReadFile(fmt.Sprintf("%s/mock/welcome.txt", dir))
	if err != nil {
		panic(err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s, %s", string(file), user))
	if _, err := e.bot.Send(msg); err != nil {
		panic(err)
	}
}
