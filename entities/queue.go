package entities

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Queue struct {
	players    []tgbotapi.User
	currentIdx int // Index of the current player's turn.
	lock       sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		players:    make([]tgbotapi.User, 0, 4), // Initializes with capacity for 4 players.
		currentIdx: -1,                          // Indicates no player's turn initially.
	}
}

// AddPlayer adds a new player to the queue.
func (q *Queue) AddPlayer(user tgbotapi.User) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// Check if user already in queue to avoid duplicates
	for _, player := range q.players {
		if player.ID == user.ID {
			return // Player already exists in the queue
		}
	}

	q.players = append(q.players, user) // Add new player
}

// StartNextTurn advances to the next player's turn.
func (q *Queue) StartNextTurn() *tgbotapi.User {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.players) == 0 {
		return nil // No players in the queue
	}

	q.currentIdx = (q.currentIdx + 1) % len(q.players) // Cycle through players
	return &q.players[q.currentIdx]
}

// GetCurrentPlayer returns the player whose turn it is.
func (q *Queue) GetCurrentPlayer() *tgbotapi.User {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.currentIdx == -1 {
		return nil // No current player's turn
	}
	return &q.players[q.currentIdx]
}
