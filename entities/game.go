package entities

import (
	"bridge/utils"
	"errors"
	"fmt"
	"math/rand"
	// "strconv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Game struct{
	Bot *tgbotapi.BotAPI
	ChatID int64
	ID uint32
	Players []*tgbotapi.User
	Deck *Deck
	// currentBid Bid
	Hands []*Hand
	InProgress bool
}

func NewGame (bot *tgbotapi.BotAPI, chatID int64) *Game{
	return &Game{
		Bot:bot,
		ChatID: chatID,
		ID: rand.Uint32(),
		Players: []*tgbotapi.User{},
		Deck : NewDeck(),
		// currentBid: Bid{},
		Hands: []*Hand{},
		InProgress: false,
	}
}

// TODO
func (g *Game) StartGame (){
	var restartGame bool
	for{
		if !g.Deck.shuffled{
			g.Deck.Shuffle()
			g.Deck.shuffled = true
			tmp := g.Deck.cards
			g.Hands = append(g.Hands, &Hand{g.Players[0],tmp[:13],make([]int,0)})
			// g.hands = append(g.hands, Hand{g.Players[1],tmp[13:26]})
			// g.hands = append(g.hands, Hand{g.Players[2],tmp[26:39]})
			// g.hands = append(g.hands, Hand{g.Players[3],tmp[39:]})
			for _,hand := range g.Hands{
				hand.SortHand()
				fmt.Println(hand.SuitIndex)
				if !hand.CountPoints(){
					g.Deck.shuffled = false
					msg:= fmt.Sprintf("Restarting Game... Player %s lacking points\n", hand.player.UserName)
					utils.SendMessage(g.Bot,g.ChatID,msg)
					restartGame = true
				}
			}

			//Check to see if need to restart game
			if !restartGame{
				//Restart game
				break
			}
			restartGame = false
		}
	}
}


func (g *Game) FindPlayer (user *tgbotapi.User) (bool,int){
	for idx,player := range g.Players{
		if player.ID == user.ID{
			return true,idx
		}
	}
	return false,-1
}

func (g *Game) AddPlayer (user *tgbotapi.User) (error){
	present,_ := g.FindPlayer(user)
	if !present {
		g.Players = append(g.Players, user)
		return nil
	}
	return errors.New("player already in game")
}


func (g *Game) RemovePlayer (user *tgbotapi.User) (error){
	present,idx := g.FindPlayer(user)
	currentPlayers := g.Players

	if present  {
		currentPlayers = append(currentPlayers[:idx], currentPlayers[idx+1:]...)
		return nil
	}
	return errors.New("player not in game")
}

func (g *Game) CheckPlayers (bot *tgbotapi.BotAPI, chatID int64, roomID uint32,msgID int){
	   if len(g.Players) == 1{
	// if len(g.Players) == 4{
		utils.DeleteButton(bot,chatID,msgID)
		// utils.SendMessage(bot,chatID,fmt.Sprintf("Starting Room %d\n\nPlayer 1: %s\nPlayer 2: %s\nPlayer 3: %s\nPlayer 4: %s",roomID,g.Players[0].UserName,g.Players[1].UserName,g.Players[2].UserName,g.Players[3].UserName))
		utils.SendMessage(bot,chatID,fmt.Sprintf("Starting Room %d\n\nPlayer 1: %s",roomID,g.Players[0].UserName))
		g.StartGame()
	} else{
		fmt.Println("Room is not full...")
	}
}

func (g *Game) GetHand (idx int) (*Hand,error){
	for index,hand := range g.Hands{
		if index==idx{
			return hand,nil
		}
	}
	return &Hand{},errors.New("Hand not found")
}









