package entities

import (
	"sort"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
type Hand struct{
	player *tgbotapi.User
	Cards []*Card
	SuitIndex []int
}

func (h *Hand) SortHand() {
	//Sort cards
	sort.Sort(BySuitAndRank(h.Cards))
	h.SuitIndex = make([]int, 0)

	//Return suits indexes
	suits:= [3]Suit{Diamonds,Hearts,Spades}
	index := 0

	for idx,card := range h.Cards{
		if len(h.SuitIndex) == 3{
			break
		}
		if card.Suit == suits[index]{
			h.SuitIndex=append(h.SuitIndex,idx)
			if index < 2{
				index = index+1
			}
		}
	}
}

type BySuitAndRank []*Card

func (a BySuitAndRank) Len() int           { return len(a) }
func (a BySuitAndRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySuitAndRank) Less(i, j int) bool {
    if a[i].Suit == a[j].Suit {
        return a[i].Rank < a[j].Rank
    }
    return a[i].Suit < a[j].Suit
}
