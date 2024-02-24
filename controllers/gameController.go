package controllers

import (
	"bridge/entities"
	"bridge/utils"
	"fmt"
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type GameController struct{
	bot *tgbotapi.BotAPI
	chatID int64
	Game *entities.Game
}

func NewGameController (bot *tgbotapi.BotAPI, chatID int64, game *entities.Game) *GameController{
	return &GameController{
		bot:bot,
		chatID:chatID,
		Game:game,
	}
}

func (gc *GameController) StartNewGame() {
	fmt.Println("Start New Game")

	Game := entities.NewGame(gc.bot,gc.chatID)
	gc.Game = Game
	room := fmt.Sprintf("join_game:%d",Game.ID)
	btn := []tgbotapi.InlineKeyboardButton{utils.CreateButton("Join Game",room)}
	keyboard := utils.CreateInlineMarkup(btn)

	// Create a message with the inline keyboard
	utils.SendMessageWithMarkup(gc.bot,gc.chatID, "Starting a new game...\nNo of Players: 0",keyboard)
}

func (gc *GameController) NotifyAddPlayer (user *tgbotapi.User, room uint32, msgID int){
	game:= gc.Game
	err:=game.AddPlayer(user)
	if err != nil{
		fmt.Println(err)
	}	else{
		var text []interface{}
		var tmp []string
		var result string
		text = append(text, fmt.Sprintf("Game Lobby %d\n\n", game.ID))
		for idx, player := range (game.Players){
			text = append(text,fmt.Sprintf("Player %d: %s\n", idx+1,player.UserName))
		}
		text = append(text, fmt.Sprintf("No of Players: %d, %d more players to start", len(game.Players),4-len(game.Players)))
		for _,item := range text{
			switch v:= item.(type){
			case string:
				tmp = append(tmp, v)
			}
		}
		result = strings.Join(tmp,"")
		room := fmt.Sprintf("join_game:%d",game.ID)
		btn := []tgbotapi.InlineKeyboardButton{utils.CreateButton("Join Game",room)}
		keyboard := utils.CreateInlineMarkup(btn)
		utils.EditMessageWithMarkup(gc.bot,gc.chatID,result,msgID,keyboard)
	}
}

func (gc *GameController) AddGame(game entities.Game) {
	gc.Game = &game
}

func (gc *GameController) RemoveGame() {
	id := gc.Game.ID
	gc.Game = &entities.Game{
		Bot:        gc.Game.Bot,
		ChatID:     gc.Game.ChatID,
		ID:         gc.Game.ID + 1, // Incrementing the ID for a new game
		Players:    []*tgbotapi.User{},
		Deck:       entities.NewDeck(), // Example: InitializeDeck creates a new deck
		Hands:      []*entities.Hand{},
		InProgress: false,
	}
	fmt.Printf("Deleted room %d\n", id)
}






