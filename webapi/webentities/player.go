package webentities

func NewPlayer() *Player {
	return &Player{}
}

func CreateNewPlayer(chatID int64, tableID uint32) *Player {
	return &Player{
		ChatID:  chatID,
		TableID: tableID,
		// Hand:     *Hand{},
		Position: 0,
		Wins:     0,
	}
}
