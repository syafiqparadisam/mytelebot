package event

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/syafiqparadisam/mytelebot/repositories"
)

type event struct {
	bot  *tgbotapi.BotAPI
	repo repositories.RepoInterface
}

func NewEvent(bot *tgbotapi.BotAPI, repo repositories.RepoInterface) *event {
	return &event{
		bot:  bot,
		repo: repo,
	}
}

func (e *event) HandleEvent(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		// check user
		var user string
		exist, err := e.repo.FindUser(update.Message.From.UserName)
		if err != nil {
			panic(err)
		}
		fmt.Println(exist)
		if len(exist) == 0 {
			result, err := e.repo.CreateUser(update.Message.From.UserName)
			user = result[0].Username
			if err != nil {
				panic(err)
			}
		} else {
			user = exist[0].Username
		}

		e.sendWelcomeMessage(update, user)
		// handle when user send message
		if update.Message != nil {
			if update.Message.Text != "" {
				e.handleMessage(update)
			}
		}

		log.Printf("Got message %s from %s\n", update.Message.Text, update.Message.From)
	}
}
