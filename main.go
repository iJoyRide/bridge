package main

import (
	"bridge/controllers"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the Telegram bot token from the environment variable
	botToken := os.Getenv("TELEGRAM_APITOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_APITOKEN not found in environment variables")
	}

	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	MessageController := controllers.NewMessageController(bot)
	MessageController.StartListening()
}
