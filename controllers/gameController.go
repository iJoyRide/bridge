package controllers

import (
	"bridge/entities"
	"bridge/utils"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var GlobalGameController GameController

type GameController struct{
	bot *tgbotapi.BotAPI
	chatID int64
	Games []*entities.Game
}

func NewGameController (bot *tgbotapi.BotAPI, chatID int64, games []*entities.Game) *GameController{
	return &GameController{
		bot:bot,
		chatID:chatID,
		Games:games,
	}
}

func (gc *GameController) StartNewGame() {
	fmt.Println("Start New Game")

	Game := entities.NewGame()
	gc.AddGame(Game)
	room := fmt.Sprintf("join_game:%d",Game.ID)
	btn := []tgbotapi.InlineKeyboardButton{utils.CreateButton("Join Game",room)}
	keyboard := utils.CreateInlineMarkup(btn)

	// Create a message with the inline keyboard
	msg := tgbotapi.NewMessage(gc.chatID, "Starting a new game...\nNo of Players: 0")
	msg.ReplyMarkup = keyboard

	// Send the message with the inline keyboard
	gc.bot.Send(msg)
}

func (gc *GameController) AddPlayer (user *tgbotapi.User, room uint64, msgID int){
	game,err := gc.GetGame(room)
	if err != nil {
		fmt.Println(err)
	} else {
		err:=game.AddPlayer(user)
		if err != nil{
			fmt.Println(err)
		}	else{
			newText := fmt.Sprintf("Starting a new game...\nNo of Players: %d", len(game.Players))
			room := fmt.Sprintf("join_game:%d",game.ID)
			btn := []tgbotapi.InlineKeyboardButton{utils.CreateButton("Join Game",room)}
			keyboard := utils.CreateInlineMarkup(btn)
			utils.EditMessageWithMarkup(gc.bot,gc.chatID,newText,msgID,&keyboard)
		}
	}
}

func (gc *GameController) GetGame (id uint64) (*entities.Game,error){
	for _,game := range gc.Games{
		if game.ID == id{
            return game,nil
		}
	}
	return nil, errors.New("no game found")
}

func (gc *GameController) AddGame(game *entities.Game) {
	gc.Games = append(gc.Games, game)
}



