package mazesolver

// MazeMeta holds meta data for the solution taken
type MazeMeta struct {
	row          int
	rowPos       int
	debug        bool
	direction    int
	directions   [4]string
	Moves        [][2]int
	IsSuccessful bool
}

func (m *MazeMeta) addMove(row int, rowPos int) {
	m.Moves = append(m.Moves, [2]int{row, rowPos})
}

func (m *MazeMeta) addMoves(moves [][2]int) {
	for _, move := range moves {
		m.addMove(move[0], move[1])
	}
}

func (m *MazeMeta) containsMove(row int, rowPos int) bool {
	for _, move := range m.Moves {
		if move[0] == row && move[1] == rowPos {
			return true
		}
	}
	return false
}
