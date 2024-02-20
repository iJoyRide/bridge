package entities

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bid struct {
	Player tgbotapi.User
	Suit   string
	Value  int
}

type BiddingSystem struct {
	Bids  []Bid
	Queue *Queue // Reference to the Queue system
}

func NewBiddingSystem(queue *Queue) *BiddingSystem {
	return &BiddingSystem{
		Bids:  make([]Bid, 0),
		Queue: queue,
	}
}

func (b *BiddingSystem) MakeBid(user tgbotapi.User, suit string, value int) error {
	if b.Queue.GetCurrentPlayer() == nil || b.Queue.GetCurrentPlayer().ID != user.ID {
		return fmt.Errorf("not the player's turn")
	}

	bid := Bid{
		Player: user,
		Suit:   suit,
		Value:  value,
	}

	b.Bids = append(b.Bids, bid)
	b.Queue.StartNextTurn() // Advance to the next player's turn after a successful bid
	return nil
}
