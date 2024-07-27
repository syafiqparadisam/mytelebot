package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/syafiqparadisam/mytelebot/event"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error reading env file from this path")
	}

	token := os.Getenv("TELEBOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	event := event.NewEvent(bot)
	event.HandleEvent(updates)
}
