package game

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// InitializeGame initializes the game and handles user joining
func InitializeGame(bot *tgbotapi.BotAPI) {
	// Initialize game state
	gameUsers := make(map[int]bool)
	gameStarted := false

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !gameStarted {
			// Handle user joining
			if update.Message.IsCommand() && update.Message.Command() == "join" {
				if len(gameUsers) < 4 {
					userID := update.Message.From.ID
					if _, ok := gameUsers[userID]; !ok {
						gameUsers[userID] = true
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "You have joined the game!"))
					} else {
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "You are already in the game!"))
					}
				} else {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "The game is full. Try again later!"))
				}
			}

			// Start the game if 4 users have joined
			if len(gameUsers) == 4 {
				gameStarted = true

			}
		}

	}
}
