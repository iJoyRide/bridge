package controllers

import (
	"bridge/entities"
	"bridge/utils"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//Class
type MessageController struct{
	//Member variables
	bot *tgbotapi.BotAPI
	GameControllers []*GameController
}

//Constructor
func NewMessageController(bot *tgbotapi.BotAPI) *MessageController{
	fmt.Println("Created Message Controller")
	return &MessageController{
		bot:     bot,
		GameControllers: []*GameController{},
	}
}

//Listener
func (mc *MessageController) StartListening() {
	// mc.bot.MakeRequest()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	fmt.Println("Start Listening...")
	updates := mc.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil{
			mc.HandleMessage(update)
			continue
		}
		if update.CallbackQuery != nil{
			mc.HandleCallbackQuery(update.CallbackQuery)
			continue
		}
		if update.InlineQuery != nil{
			err:=mc.HandleInlineQuery(update.InlineQuery)
			if err!=nil{
				log.Println(err)
			}else{
				continue
			}
		}
		// if update.Message == nil{
		// 	if update.CallbackQuery !=nil{
		// 		mc.HandleCallbackQuery(update.CallbackQuery)
		// 	} else{
		// 		continue
		// 	}
		// } else{
		// 	mc.HandleMessage(update)
		// }
	}
}

func (mc *MessageController) CheckGameController (gc *GameController) bool{
	if len(mc.GameControllers) == 0{
		return false //dont exist
	}
	for _,c := range mc.GameControllers{
		if c == gc{
			return true
		}
	}
	return false
}

func (mc *MessageController) AddGameController (gc *GameController){
	if !mc.CheckGameController(gc){
		mc.GameControllers = append(mc.GameControllers, gc)
		fmt.Println("Added game controller to list!")
		return
	} else{
		fmt.Printf("From existing game controller: chat %d\n", gc.chatID)
		return
	}
}

func (mc *MessageController)CheckOngoingController(chatID int64) bool{
	_,err:=mc.FindGameController(chatID)
	if err!=nil{
		fmt.Println(err)
		return false
	}else{
		return true
	}
}
//MessageHandler
func (mc *MessageController) HandleMessage(update tgbotapi.Update) {
	if update.Message.IsCommand(){
		command := update.Message.Command()
		switch command {
		case "start":
			utils.SendMessage(mc.bot,update.Message.Chat.ID, "Welcome to Bridge! Bridge is a four-player partnership trick-taking game with thirteen tricks per deal\n\n/help - For more commands")
		case "help":
			utils.SendMessage(mc.bot,update.Message.Chat.ID, "Available commands:\n/start - Start the bot\n/play_game - Start a new game\n/leave - Leave game")
		case "play_game":
			// Check game controller
			if !mc.CheckOngoingController(update.Message.Chat.ID){
				gameController:=GameController{mc.bot,update.Message.Chat.ID,entities.NewGame(mc.bot,update.Message.Chat.ID)}
				mc.AddGameController(&gameController)
				mc.PrintAllControllers()
				gameController.StartNewGame()
			}else{
				utils.SendMessage(mc.bot,update.Message.Chat.ID,fmt.Sprintf("%s, a game is already ongoing...",update.Message.From.UserName))
			}
		case "leave":
			gc,err := mc.FindGameController(update.Message.Chat.ID)
			if err != nil{
				fmt.Println(err)
			}else{
				gc.Game.RemovePlayer(update.Message.From)
				msg := fmt.Sprintf("%s has left room %d\n\nShutting down game...", update.Message.From, gc.Game.ID)
				utils.SendMessage(mc.bot,update.Message.Chat.ID,msg)
				gc.RemoveGame()
				mc.RemoveGameController(update.Message.Chat.ID)
			}
		default:
			utils.SendMessage(mc.bot,update.Message.Chat.ID, "Unknown command. Type /help for a list of available commands.")
		}
	}else if update.Message.Sticker != nil{
		sticker := update.Message.Sticker
		chatID := update.Message.Chat.ID
		user := update.Message.From
		id := entities.IDToName(sticker.FileUniqueID)
		gc,err := mc.FindGameController(chatID)
		if !gc.Game.InProgress{
			if err != nil{
				fmt.Println(err)
			}else{
				game := gc.Game
				_,idx:=game.FindPlayer(user)
				hand := gc.Game.Hands[idx]
				hand.RemoveCard(id)
			}
		}else{
			fmt.Println("Game has not started, cannot throw card")
		}
	}
}

//Callback Query Handler
func (mc *MessageController) HandleCallbackQuery (query *tgbotapi.CallbackQuery) {
	// Extract relevant information from the callback query
	user := query.From
	msgID := query.Message.MessageID
	parts := strings.Split(query.Data,":") //Split by ":", use different characters for subsequent splits
	command := parts[0]
	data := parts[1]

	// Handle the callback query logic based on the data
	switch command {
	case "join_game":
		// Respond to the button click
		roomID,err:=strconv.ParseUint(data,10,32)
		duplicate,id := mc.CheckPlayerDuplicate(query.Message.Chat.ID,user)
		if duplicate && id != 0{ //Check if player is in another game
			fmt.Println("User in another room, not allowed")
			msg := fmt.Sprintf("Player %s already in room, leave that room to join this game.", user.UserName)
			btn := []tgbotapi.InlineKeyboardButton{utils.CreateButton("Force Quit","quit_game:"+strconv.FormatInt(id,10)+"_"+strconv.FormatInt(user.ID,10))}
			keyboard := utils.CreateInlineMarkup(btn)
			utils.SendMessageWithMarkup(mc.bot,query.Message.Chat.ID,msg,keyboard)
		}else{
			fmt.Printf("Room ID: %d, user: %s pressed the button.\n",roomID, user.UserName)
			if err != nil {
				fmt.Println(err)
			} else {
				roomID:=uint32(roomID)
				gc,err := mc.FindGameController(query.Message.Chat.ID)
				if err != nil{
					fmt.Println(err)
				}else{
					game := gc.Game
					// if len(game.Players) < 4{
					if len(game.Players) < 1{
							gc.NotifyAddPlayer(query.From,roomID,msgID) //Add Player
							game:= gc.Game
							game.CheckPlayers(mc.bot,query.Message.Chat.ID,roomID,msgID) //Check if room is full, else start game
					}
				}
			}
		}
	case "quit_game": //In the case where group/chat is deleted and user unable to leave game, allow them to force quit from current chat
		quit_game_split := strings.Split(data,"_") //Split by "_"
		chatID,_ := strconv.ParseInt(quit_game_split[0],10,64)
		userID,_ := strconv.ParseInt(quit_game_split[1],10,64)
		if query.From.ID == userID{
			gc,err := mc.FindGameController(chatID)
			if err!=nil{
				fmt.Println(err)
			}else{
				id := gc.chatID
				mc.RemoveGameController(id)
				gc.RemoveGame()
				fmt.Printf("Removed game controller %d\n", id)
				utils.SendMessage(mc.bot,chatID,fmt.Sprintf("Someone left the game"))
			}
		}
	default:
		// Handle other callback query scenarios
	}
}

func (mc *MessageController) HandleInlineQuery (query *tgbotapi.InlineQuery) error{
	//Get User's cards
	user := query.From
	currentGC,err := mc.FindGameController(user)
	if currentGC == nil || currentGC.Game == nil{
		return errors.New("handleInlineQuery: no game")
	}else{
		if err != nil{
			log.Println(err)
		}else{
			_,playerIdx := currentGC.Game.FindPlayer(user)
			playerHand,err := currentGC.Game.GetHand(playerIdx)
			if err!= nil{
				log.Println(err)
			}else{
				var stickers []interface{}
				for idx,card:= range playerHand.Cards{
					id,err := strconv.Atoi(query.ID)
					if err!= nil{
						log.Println(err)
					}else{
						c := fmt.Sprintf("%s_%d", card.Suit, card.Rank)
						//Search for card ID
						cardID := entities.NameToID(c)
						article := tgbotapi.NewInlineQueryResultCachedSticker(strconv.Itoa(id+idx),cardID,c) //nto sure if rand is best choice
						stickers = append(stickers,article)
					}
				}

				inlineConfig := tgbotapi.InlineConfig{
					InlineQueryID: query.ID,
					IsPersonal: true,
					CacheTime: 1,
					Results: stickers,
				}

				_, err := mc.bot.Request(inlineConfig)
				if err != nil {
					fmt.Println("Error answering inline query:", err)
				}
			}
		}
	}
	return nil

}

func (mc *MessageController) CheckPlayerDuplicate (chatID int64, user *tgbotapi.User) (bool,int64){
	for _, gc := range mc.GameControllers{
		inGame,_:=gc.Game.FindPlayer(user)
		if inGame && gc.Game.ChatID != chatID{
			return true,gc.Game.ChatID
		}
	}

	return false,0
}

func (mc *MessageController) FindGameController (e interface{}) (*GameController,error){
	switch m := e.(type) {
	case int64: //chatID
		for _,controller := range mc.GameControllers{
			if controller.chatID == m{
				return controller,nil
			}
		}
	case *tgbotapi.User:
		for _,controller := range mc.GameControllers{
				for _,player := range controller.Game.Players{
					if player.ID == m.ID{
						return controller,nil
					}
				}
			}
	}
	return nil,errors.New("no controller found/user not found")
}

func (mc *MessageController) RemoveGameController (chatID int64){
	var index int
	for idx,controller := range mc.GameControllers{
		if controller.chatID == chatID{
			index = idx
			break
		}
	}
	mc.GameControllers = append(mc.GameControllers[:index],mc.GameControllers[index+1:]...)
}

func (mc *MessageController) PrintAllControllers (){
	for _,controller := range mc.GameControllers{
		fmt.Printf("ChatID: %d\n", controller.chatID)
	}
}