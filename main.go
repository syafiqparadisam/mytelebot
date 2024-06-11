package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	token := os.Getenv("TELEBOT_TOKEN")
	tgbotapi.NewBotAPI(token)
}