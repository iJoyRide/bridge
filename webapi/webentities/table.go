package webentities

func NewTable() *Table {
	return &Table{}
}

func CreateNewTable(tableID uint32) *Table {
	return &Table{
		TableID: tableID,
		Count:   1,
		Trump:   "",
		Bid:     0,
		Queue:   0,
		Team_A:  make([]int16, 2),
		A_Score: 0,
		Team_B:  make([]int16, 2),
		B_Score: 0,
	}
}
