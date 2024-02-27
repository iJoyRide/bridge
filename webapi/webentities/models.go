package webentities

// type Game struct {
// 	Bot        *tgbotapi.BotAPI `bson:"bot"`
// 	ChatID     int64            `bson:"chat_id"`
// 	ID         uint32           `bson:"id"`
// 	Players    []*tgbotapi.User `bson:"players"`
// 	Deck       *Deck            `bson:"deck"`
// 	Hands      []*Hand          `bson:"hands"`
// 	InProgress bool             `bson:"in_progress"`
// }

type Player struct {
	ChatID   int64  `bson:"chat_id"`
	TableID  uint32 `bson:"table_id"`
	Hand     *Hand  `bson:"hand"`
	Position int16  `bson:"position"`
	Wins     int16  `bson:"wins"`
}

type Table struct {
	TableID uint32  `bson:"table_id"`
	Count   int16   `bson:"count"`
	Trump   string  `bson:"trump"`
	Bid     int16   `bson:"bid"`
	Queue   int16   `bson:"queue"`
	Team_A  []int16 `bson:"team_a"`
	A_Score int16   `bson:"a_score"`
	Team_B  []int16 `bson:"team_b"`
	B_Score int16   `bson:"b_score"`
}

type Hand struct {
	Cards     []*Card `bson:"cards"`
	SuitIndex []int   `bson:"suit_index"`
}

type Card struct {
	Suit Suit `bson:"suit"`
	Rank Rank `bson:"rank"`
}

type Suit string

const (
	Spades   Suit = "Spades"
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
)

type Deck struct {
	cards    []*Card `bson:"cards"`
	shuffled bool    `bson:"shuffled"`
}

type Rank int

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
