package webentities

import (
	"fmt"
	"math/rand"
)

func NewDeck() *Deck {
	cards := LoadDeck()
	return &Deck{
		cards:    cards,
		shuffled: false,
	}
}

func LoadDeck() []*Card {
	var cards []*Card
	fmt.Println("Loading Deck for new game...")

	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for rank := Two; rank <= Ace; rank++ {
			cards = append(cards, &Card{suit, rank})
		}
	}

	return cards
}

func (deck *Deck) Shuffle() {
	rand.Shuffle(len(deck.cards), func(i, j int) { deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i] })
}

func (deck *Deck) Deal() {
}
