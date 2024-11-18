package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
	"github.com/syafiqparadisam/mytelebot/event"
	"github.com/syafiqparadisam/mytelebot/repositories"
	// "github.com/supabase-community/supabase-go"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error reading env file from this path")
	}

	// connectdb
	apiUrl := os.Getenv("API_URL")
	apiKey := os.Getenv("API_KEY")
	client, err := supabase.NewClient(apiUrl, apiKey, nil)
	if err != nil {
		panic(err)
	}

	repo := repositories.NewRepository(client)

	token := os.Getenv("TELEBOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	event := event.NewEvent(bot, repo)
	event.HandleEvent(updates)
}
