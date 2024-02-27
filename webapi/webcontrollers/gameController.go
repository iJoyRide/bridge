package webcontrollers

import (
	"fmt"

	"bridge/utils"
	"bridge/webapi/webentities"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type GameController struct {
	Player webentities.Player
	Table  webentities.Table
}

func NewGameController() *GameController {
	fmt.Println("Created Empty Game Controller")
	var player webentities.Player
	var table webentities.Table
	return &GameController{
		Player: player,
		Table:  table,
	}
}

func CreateGameController(chatID int64, tableID uint32) *GameController {
	fmt.Println("Created Message Controller")
	return &GameController{
		Player: *webentities.CreateNewPlayer(chatID, tableID),
		Table:  *webentities.CreateNewTable(tableID),
	}
}

func (gc *GameController) StartNewGame(bot *tgbotapi.BotAPI) {
	fmt.Println("Start New Game")
	room := fmt.Sprintf("join_game:%d", gc.Player.TableID)
	btn := []tgbotapi.InlineKeyboardButton{utils.CreateButton("Join Game", room)}
	keyboard := utils.CreateInlineMarkup(btn)

	// Create a message with the inline keyboard
	utils.SendMessageWithMarkup(bot, gc.Player.ChatID, "Starting a new game...\nNo of Players: 0", keyboard)
}
