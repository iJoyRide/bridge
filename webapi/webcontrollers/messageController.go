package webcontrollers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"bridge/utils"
	"bridge/webapi/webentities"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// Class
type MessageController struct {
	Bot  *tgbotapi.BotAPI
	Game *GameController
	DB   *DatabaseController
}

func SetUpBot() *tgbotapi.BotAPI {
	Bot := &tgbotapi.BotAPI{
		Token:  os.Getenv("BOT_TOKEN"),
		Client: &http.Client{},
		Buffer: 100,
	}
	Bot.SetAPIEndpoint(tgbotapi.APIEndpoint)
	return Bot
}

// Constructor
func NewMessageController(bot *tgbotapi.BotAPI, db *mongo.Database) *MessageController {
	fmt.Println("Created Message Controller")
	return &MessageController{
		Bot:  bot,
		Game: NewGameController(),
		DB:   NewDatabaseController(db),
	}
}

func (mc *MessageController) StartListening(w http.ResponseWriter, r *http.Request) {

	updates := mc.Bot.ListenForWebhookRespReqFormat(w, r)
	fmt.Println("Start Listening...")

	for update := range updates {
		if update.Message != nil {
			mc.HandleMessage(update)
			continue
		}
		if update.CallbackQuery != nil {
			mc.HandleCallbackQuery(update.CallbackQuery)
			continue
		}
		// if update.InlineQuery != nil {
		// 	err := mc.HandleInlineQuery(update.InlineQuery)
		// 	if err != nil {
		// 		log.Println(err)
		// 	} else {
		// 		continue
		// 	}
		// }

	}
}

func (mc *MessageController) HandleMessage(update tgbotapi.Update) {
	if update.Message.IsCommand() {
		command := update.Message.Command()
		switch command {
		case "start":
			utils.SendMessage(mc.Bot, update.Message.Chat.ID, "Welcome to Bridge! Bridge is a four-player partnership trick-taking game with thirteen tricks per deal\n\n/help - For more commands")

		case "help":
			utils.SendMessage(mc.Bot, update.Message.Chat.ID, "Available commands:\n/start - Start the bot\n/play_game - Start a new game\n/leave - Leave game")

		case "play_game":
			utils.SendMessage(mc.Bot, update.Message.Chat.ID, "A Player is starting game")
			if !mc.CheckPlayer(update.Message.Chat.ID) {
				tableID, chatID := rand.Uint32(), update.Message.Chat.ID
				mc.Game = CreateGameController(chatID, tableID)
				mc.PrintAll()
				mc.Game.StartNewGame(mc.Bot)
				mc.DB.InsertPlayer(&mc.Game.Player)
				mc.DB.InsertTable(&mc.Game.Table)

			} else {
				utils.SendMessage(mc.Bot, update.Message.Chat.ID, fmt.Sprintf("%s, a game is already ongoing...", update.Message.From.UserName))
			}

		case "leave":
			utils.SendMessage(mc.Bot, update.Message.Chat.ID, "A Player is leaving game")
			if mc.CheckPlayer(update.Message.Chat.ID) {
				mc.DB.GetPlayerByChatID(update.Message.Chat.ID, &mc.Game.Player)
				mc.DB.DeletePlayersByTableID(*&mc.Game.Player.TableID)
				msg := fmt.Sprintf("%s has left room %d\n\nShutting down game...", update.Message.From, mc.Game.Player.TableID)
				utils.SendMessage(mc.Bot, update.Message.Chat.ID, msg)
			} else {
				msg := fmt.Sprintf("%s has no room to exit", update.Message.From)
				utils.SendMessage(mc.Bot, update.Message.Chat.ID, msg)
			}

			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	gc.Game.RemovePlayer(update.Message.From)
			// 	msg := fmt.Sprintf("%s has left room %d\n\nShutting down game...", update.Message.From, gc.Game.ID)
			// 	utils.SendMessage(mc.Bot, update.Message.Chat.ID, msg)
			// 	gc.RemoveGame()
			// 	mc.RemoveGameController(update.Message.Chat.ID)
			// }
		default:
			utils.SendMessage(mc.Bot, update.Message.Chat.ID, "Unknown command. Type /help for a list of available commands.")
		}
	}

}

func (mc *MessageController) CheckPlayer(chatID int64) bool {
	err := mc.DB.GetPlayerByChatID(chatID, &mc.Game.Player)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func (mc *MessageController) PrintAll() {
	messageControllerJSON, err := json.Marshal(mc)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	// Print the JSON representation
	fmt.Println(string(messageControllerJSON))

}

func (mc *MessageController) HandleCallbackQuery(query *tgbotapi.CallbackQuery) {
	user := query.From
	// msgID := query.Message.MessageID
	parts := strings.Split(query.Data, ":") //Split by ":", use different characters for subsequent splits
	command := parts[0]
	data := parts[1]
	switch command {
	case "join_game":
		roomID, _ := strconv.ParseUint(data, 10, 32)
		if !mc.CheckPlayer(user.ID) {
			mc.Game.Player = *webentities.CreateNewPlayer(user.ID, uint32(roomID))
			mc.DB.GetTableByTableID(uint32(roomID), &mc.Game.Table)
			mc.PrintAll()
			mc.Game.StartNewGame(mc.Bot)
			mc.DB.InsertPlayer(&mc.Game.Player)
			mc.DB.UpdateTable(&mc.Game.Table)
		}
		// } else {
		// 	fmt.Println("User in another room, not allowed")
		// 	msg := fmt.Sprintf("Player %s already in room, leave that room to join this game.", user.UserName)
		// 	btn := []tgbotapi.InlineKeyboardButton{utils.CreateButton("Force Quit", "quit_game:"+strconv.FormatInt(int64(roomID), 10)+"_"+strconv.FormatInt(user.ID, 10))}
		// 	keyboard := utils.CreateInlineMarkup(btn)
		// 	utils.SendMessageWithMarkup(mc.Bot, query.Message.Chat.ID, msg, keyboard)
		// }

	// case "quit_game": //In the case where group/chat is deleted and user unable to leave game, allow them to force quit from current chat

	default:
		// Handle other callback query scenarios
	}

}

// 	// Extract relevant information from the callback query
// 	user := query.From
// 	msgID := query.Message.MessageID
// 	parts := strings.Split(query.Data, ":") //Split by ":", use different characters for subsequent splits
// 	command := parts[0]
// 	data := parts[1]

// 	// Handle the callback query logic based on the data
// 	switch command {
// 	case "join_game":
// 		// Respond to the button click
// 		roomID, err := strconv.ParseUint(data, 10, 32)
// 		duplicate, id := mc.CheckPlayerDuplicate(query.Message.Chat.ID, user)
// 		if duplicate && id != 0 { //Check if player is in another game
// 			fmt.Println("User in another room, not allowed")
// 			msg := fmt.Sprintf("Player %s already in room, leave that room to join this game.", user.UserName)
// 			btn := []tgbotapi.InlineKeyboardButton{utils.CreateButton("Force Quit", "quit_game:"+strconv.FormatInt(id, 10)+"_"+strconv.FormatInt(user.ID, 10))}
// 			keyboard := utils.CreateInlineMarkup(btn)
// 			utils.SendMessageWithMarkup(mc.bot, query.Message.Chat.ID, msg, keyboard)
// 		} else {
// 			fmt.Printf("Room ID: %d, user: %s pressed the button.\n", roomID, user.UserName)
// 			if err != nil {
// 				fmt.Println(err)
// 			} else {
// 				roomID := uint32(roomID)
// 				gc, err := mc.FindGameController(query.Message.Chat.ID)
// 				if err != nil {
// 					fmt.Println(err)
// 				} else {
// 					game := gc.Game
// 					// if len(game.Players) < 4{
// 					if len(game.Players) < 1 {
// 						gc.NotifyAddPlayer(query.From, roomID, msgID) //Add Player
// 						game := gc.Game
// 						game.CheckPlayers(mc.bot, query.Message.Chat.ID, roomID, msgID) //Check if room is full, else start game
// 					}
// 				}
// 			}
// 		}
// 	case "quit_game": //In the case where group/chat is deleted and user unable to leave game, allow them to force quit from current chat
// 		quit_game_split := strings.Split(data, "_") //Split by "_"
// 		chatID, _ := strconv.ParseInt(quit_game_split[0], 10, 64)
// 		userID, _ := strconv.ParseInt(quit_game_split[1], 10, 64)
// 		if query.From.ID == userID {
// 			gc, err := mc.FindGameController(chatID)
// 			if err != nil {
// 				fmt.Println(err)
// 			} else {
// 				id := gc.chatID
// 				mc.RemoveGameController(id)
// 				gc.RemoveGame()
// 				fmt.Printf("Removed game controller %d\n", id)
// 				utils.SendMessage(mc.Bot, chatID, fmt.Sprintf("Someone left the game"))
// 			}
// 		}
// 	default:
// 		// Handle other callback query scenarios
// 	}
// }
