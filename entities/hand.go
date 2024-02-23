package entities

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
type Hand struct{
	player *tgbotapi.User
	Cards []*Card
	SuitIndex []int
}

func (h *Hand) RemoveCard(id string){
	//Find card
	var index int
	index = -1
	parts := strings.Split(id,"_")
	suit := parts[0]
	rank,err := strconv.Atoi(parts[1])
	if err!=nil{
		fmt.Println("cannot convert to str:RemoveCard")
	}
	//loop through hand
	for idx,card := range h.Cards{
		if suit == string(card.Suit) && rank == int(card.Rank){
			index =idx
			break
		}
	}

	if index == -1{
		fmt.Printf("Player %s does not have card %s\n", h.player, id)
	}else{
		//Remove card
		h.Cards = append(h.Cards[:index], h.Cards[index+1:]...)

		//Resort card
		h.SortHand()
	}
}

func (h *Hand) CountPoints() bool{
	//Count no of suits
	var points int
	var index []int
	points = 0
	index = append(index, h.SuitIndex[0])
	index = append(index,h.SuitIndex[1]-h.SuitIndex[0])
	index = append(index,h.SuitIndex[2]-h.SuitIndex[1])
	index = append(index,13 - h.SuitIndex[2])

	for _,suit := range index{
		if suit >= 5{
			points += suit-4
		}
	}

	for _,card := range h.Cards{
		if int(card.Rank) > 10{
			points += int(card.Rank)-10
		}
	}

	return points >= 4
}

func (h *Hand) SortHand() {
	//Sort cards
	sort.Sort(BySuitAndRank(h.Cards))
	h.SuitIndex = make([]int, 0)

	//Count points

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
