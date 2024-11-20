package event

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/syafiqparadisam/mytelebot/entity"
	"github.com/syafiqparadisam/mytelebot/repositories"
)

type event struct {
	bot    *tgbotapi.BotAPI
	repo   repositories.RepoInterface
	chatId int64
	chat   string
	user   string
}

func NewEvent(bot *tgbotapi.BotAPI, repo repositories.RepoInterface) *event {
	return &event{
		bot:    bot,
		repo:   repo,
		chatId: 1,
	}
}

func (e *event) HandleEvent(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		chats := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		e.chatId = chats.ChatID
		e.chat = chats.Text

		// check user
		userFromDb, err := e.repo.FindUser(update.Message.From.UserName)
		if err != nil {
			panic(err)
		}

		users := *userFromDb
		if len(users) == 0 {
			userEntity := &entity.User{
				Username:  update.Message.From.UserName,
				Lastname:  update.Message.From.LastName,
				Firstname: update.Message.From.FirstName,
				ChatId:    e.chatId,
			}

			err := e.repo.CreateUser(userEntity)
			if err != nil {
				panic(err)
			}

			e.user = userEntity.Username
		} else {
			e.user = users[0].Username
		}

		e.handleMessage()

		log.Printf("Got message %s from %s\n", e.chat, update.Message.From)
	}
}

// // Membuat keyboard dinamis
// func (e *event) updateKeyboard() tgbotapi.InlineKeyboardMarkup {
// 	inlineKeys := [][]tgbotapi.InlineKeyboardButton{
// 		{
// 			tgbotapi.NewInlineKeyboardButtonData(e.GetCheckboxTest("Option 1", checkboxState["option1"]), "option1"),
// 			tgbotapi.NewInlineKeyboardButtonData(e.GetCheckboxTest("Option 2", checkboxState["option2"]), "option2"),
// 		},
// 		{
// 			tgbotapi.NewInlineKeyboardButtonData(e.GetCheckboxTest("Option 3", checkboxState["option3"]), "option3"),
// 		},
// 		{
// 			tgbotapi.NewInlineKeyboardButtonData("Submit", "submit"),
// 		},
// 	}

// 	return tgbotapi.NewInlineKeyboardMarkup(inlineKeys...)
// }

// Membuat teks checkbox berdasarkan status
func (e *event) GetCheckboxTest(label string, checked bool) string {
	if checked {
		return "✅ " + label
	}
	return "☑️ " + label
}
