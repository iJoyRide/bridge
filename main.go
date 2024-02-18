package main

import (
	game "github.com/iJoyRide/bridge/init"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("TELEGRAM_APITOKEN")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	game.InitializeGame(bot)
}
