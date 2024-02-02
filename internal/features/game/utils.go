package game

func (g *TicTacToeGame) isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}
