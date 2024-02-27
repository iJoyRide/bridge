package webcontrollers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"bridge/utils"

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
		// if update.CallbackQuery != nil {
		// 	mc.HandleCallbackQuery(update.CallbackQuery)
		// 	continue
		// }
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
			// gc, err := mc.FindGameController(update.Message.Chat.ID)
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
