package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Suit string
type Rank int

const (
	Spades   Suit = "Spades"
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
)

const (
	Ace   Rank = 14
	King  Rank = 13
	Queen Rank = 12
	Jack  Rank = 11
	Ten   Rank = 10
	Nine  Rank = 9
	Eight Rank = 8
	Seven Rank = 7
	Six   Rank = 6
	Five  Rank = 5
	Four  Rank = 4
	Three Rank = 3
	Two   Rank = 2
)

type Card struct {
	Suit Suit
	Rank Rank
}

type Deck []Card
type Hand []Card

type GameSession struct {
	Players []tgbotapi.User
	Deck    Deck
	Hands   []Hand
}

var gameInProgress bool
var session GameSession

func sortCards(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Suit != cards[j].Suit {
			return cards[i].Suit < cards[j].Suit
		}
		return cards[i].Rank > cards[j].Rank
	})
}

func NewDeck() Deck {
	ranks := []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}

	var deck Deck
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

func Shuffle(deck Deck) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
}

func Deal(deck Deck, numPlayers int) []Hand {
	hands := make([]Hand, numPlayers)
	for i := 0; i < numPlayers; i++ {
		hands[i] = make(Hand, 0, len(deck)/numPlayers)
	}

	for i, card := range deck {
		hands[i%numPlayers] = append(hands[i%numPlayers], card)
	}

	return hands
}

func suitEmoji(suit Suit) string {
	switch suit {
	case Spades:
		return "♠️"
	case Hearts:
		return "❤️"
	case Diamonds:
		return "♦️"
	case Clubs:
		return "♣️"
	default:
		return ""
	}
}

func rankEmoji(rank Rank) string {
	switch rank {
	case Ace:
		return "A"
	case King:
		return "K"
	case Queen:
		return "Q"
	case Jack:
		return "J"
	default:
		return fmt.Sprintf("%d", rank)
	}
}

func promptPartnerSelection(bot *tgbotapi.BotAPI, winnerUserID int64) {
	// Prompt the user to select a suit
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	var rows [][]tgbotapi.InlineKeyboardButton
	for _, suit := range suits {
		buttonText := fmt.Sprintf("Select %s", suitEmoji(suit))
		callbackData := fmt.Sprintf("select_partner_suit_%s", suit)
		button := tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackData)
		row := []tgbotapi.InlineKeyboardButton{button}
		rows = append(rows, row)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
	msg := tgbotapi.NewMessage(winnerUserID, "Select the suit for your partner:")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func promptRankSelection(bot *tgbotapi.BotAPI, winnerUserID int64, chosenSuit Suit) {
	// Prompt the user to select a rank
	ranks := []Rank{Ten, Jack, Queen, King, Ace}
	var rows [][]tgbotapi.InlineKeyboardButton
	for _, rank := range ranks {
		buttonText := fmt.Sprintf("Select %s", rankEmoji(rank))
		callbackData := fmt.Sprintf("select_partner_rank_%s_%d", chosenSuit, rank)
		button := tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackData)
		row := []tgbotapi.InlineKeyboardButton{button}
		rows = append(rows, row)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
	msg := tgbotapi.NewMessage(winnerUserID, "Select the rank for your partner:")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func startGame(bot *tgbotapi.BotAPI, chatID int64) {
	if !gameInProgress {
		log.Printf("Game session started")
		gameInProgress = true

		session.Deck = NewDeck()
		Shuffle(session.Deck)
		session.Hands = Deal(session.Deck, len(session.Players))

		for i := range session.Players {
			playerHand := session.Hands[i]
			sortCards(playerHand)
			var rows [][]tgbotapi.InlineKeyboardButton
			for _, card := range playerHand {
				buttonText := fmt.Sprintf("%s%s", suitEmoji(card.Suit), rankEmoji(card.Rank))
				callbackData := fmt.Sprintf("card_%s_%d", card.Suit, card.Rank)
				button := tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackData)
				row := []tgbotapi.InlineKeyboardButton{button}
				rows = append(rows, row)
			}
			keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
			msg := tgbotapi.NewMessage(chatID, "Your hand:")
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		}
		// Prompt the winner to select the partner's suit
		promptPartnerSelection(bot, int64(session.Players[0].ID))
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("TELEGRAM_APITOKEN")
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates: %v", err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			callbackData := update.CallbackQuery.Data

			if strings.HasPrefix(callbackData, "select_partner_suit_") {
				selectedSuit := strings.TrimPrefix(callbackData, "select_partner_suit_")

				var suit Suit
				switch selectedSuit {
				case "Spades":
					suit = Spades
				case "Hearts":
					suit = Hearts
				case "Diamonds":
					suit = Diamonds
				case "Clubs":
					suit = Clubs
				}

				promptRankSelection(bot, int64(update.CallbackQuery.From.ID), suit)
			} else if strings.HasPrefix(callbackData, "select_partner_rank_") {
				parts := strings.Split(callbackData, "_")
				if len(parts) == 5 {
					selectedSuit := parts[3]
					selectedRank := parts[4]

					rankInt, err := strconv.Atoi(selectedRank)
					if err != nil {
						log.Printf("Error converting rank to integer: %v", err)
					} else {
						selectedRankType := Rank(rankInt)

						confirmationMsg := fmt.Sprintf("The partner is  %s of %s.",
							rankEmoji(selectedRankType), suitEmoji(Suit(selectedSuit)))
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, confirmationMsg)
						if _, err := bot.Send(msg); err != nil {
							log.Printf("Error sending confirmation message: %v", err)
						}
					}
				}
			}

			if _, err := bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{
				CallbackQueryID: update.CallbackQuery.ID,
			}); err != nil {
				log.Printf("Error clearing callback query: %v", err)
			}
		} else if update.Message != nil {
			command := strings.Split(update.Message.Text, " ")[0]

			switch command {
			case "/start":
				msgText := "Welcome to the Bridge bot! Use /help to see available commands."
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
				_, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}

				// Sending a separate message to prompt the user to start the game
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("🎲 Play Bridge", "play_game"),
					),
				)
				startMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Press 'Play Bridge' to start the game.")
				startMsg.ReplyMarkup = keyboard
				_, err = bot.Send(startMsg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}

			case "/help":
				msgText := "Available commands:\n" +
					"/start - Start interacting with the bot\n" +
					"/help - Get help and see available commands\n" +
					"/play_game - Start a new game of Bridge"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
				_, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			case "/play_game":
				if !gameInProgress {
					// Create a new session if none exists
					if session.Players == nil {
						session.Players = make([]tgbotapi.User, 0)
					}
					// Add player to the session
					session.Players = append(session.Players, *update.Message.From)

					if len(session.Players) == 4 {
						startGame(bot, update.Message.Chat.ID)
					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Waiting for players to join... Current players: "+fmt.Sprint(len(session.Players)))
						_, err := bot.Send(msg)
						if err != nil {
							log.Printf("Error sending message: %v", err)
						}
					}
				}
			case "/leave":
				if gameInProgress {
					// Logic to handle leaving the game
					gameInProgress = false
					session = GameSession{} // Clear session data
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You have left the game.")
					_, err := bot.Send(msg)
					if err != nil {
						log.Printf("Error sending message: %v", err)
					}
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is no game in progress.")
					_, err := bot.Send(msg)
					if err != nil {
						log.Printf("Error sending message: %v", err)
					}
				}
			default:
				msgText := "Sorry, I didn't recognize that command. Use /help to see available commands."
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
				_, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			}
		}
	}
}
