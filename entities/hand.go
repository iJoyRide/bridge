package entities
import(
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)
type Hand struct{
	player *tgbotapi.User
	Cards []Card
}

func NewHand (player *tgbotapi.User, cards []Card) *Hand{
	return &Hand{
		player:player,
		Cards:cards,
	}
}

