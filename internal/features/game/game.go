package game

import (
	"errors"
	"sync"
)

type Player string

const (
	PlayerX Player = "X"
	PlayerO Player = "O"
	Empty   Player = " "
)

type TicTacToeGame struct {
	Board     [3][3]Player
	Turn      Player
	Winner    Player
	PlayerXID string
	PlayerOID string
	Mutex     sync.Mutex
}

func NewTicTacToeGame() *TicTacToeGame {
	return &TicTacToeGame{
		Turn:   PlayerX,
		Winner: Empty,
		Board: [3][3]Player{
			{Empty, Empty, Empty},
			{Empty, Empty, Empty},
			{Empty, Empty, Empty},
		},
	}
}

func (g *TicTacToeGame) PlayMove(x, y int, authorID string) error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	if (g.Turn == PlayerX && authorID != g.PlayerXID) || (g.Turn == PlayerO && authorID != g.PlayerOID) {
		return errors.New("it's not your turn")
	}
	if x < 0 || y < 0 || x >= 3 || y >= 3 {
		return errors.New("invalid move")
	}
	if g.Board[x][y] != Empty {
		return errors.New("cell is already occupied")
	}
	if g.Winner != Empty {
		return errors.New("game has already ended")
	}

	g.Board[x][y] = g.Turn
	if g.checkWin() {
		g.Winner = g.Turn
	} else {
		g.Turn = g.switchTurn()
	}

	return nil
}

func (g *TicTacToeGame) switchTurn() Player {
	if g.Turn == PlayerX {
		return PlayerO
	}
	return PlayerX
}

func (g *TicTacToeGame) checkWin() bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.Board[i][0] == g.Board[i][1] && g.Board[i][1] == g.Board[i][2] && g.Board[i][0] != Empty {
			return true
		}
	}
	// Check columns
	for j := 0; j < 3; j++ {
		if g.Board[0][j] == g.Board[1][j] && g.Board[1][j] == g.Board[2][j] && g.Board[0][j] != Empty {
			return true
		}
	}
	// Check diagonals
	if g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] && g.Board[0][0] != Empty {
		return true
	}
	if g.Board[0][2] == g.Board[1][1] && g.Board[1][1] == g.Board[2][0] && g.Board[2][0] != Empty {
		return true
	}

	return false
}

func (g *TicTacToeGame) BoardString() string {
	var boardStr string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boardStr += string(g.Board[i][j]) + " "
		}
		boardStr += "\n"
	}
	return boardStr
}
