package game

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type TicTacToeCommand struct {
	ActiveGames map[string]*TicTacToeGame
	Mutex       sync.Mutex
}

func (cmd *TicTacToeCommand) Description() string {
	return "Play Tic Tac Toe! Start a game with `-tictactoe start`, make moves with `-tictactoe x y`."
}

func NewTicTacToeCommand() *TicTacToeCommand {
	return &TicTacToeCommand{
		ActiveGames: make(map[string]*TicTacToeGame),
	}
}

func (cmd *TicTacToeCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	//args := strings.Fields(m.Content)

	// Ensure there's at least a command and one argument
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Usage: `-tictactoe start` to begin a new game or `-tictactoe x y` to make a move.")
		return
	}

	cmd.Mutex.Lock()
	defer cmd.Mutex.Unlock()

	// Handle starting a new game
	if args[1] == "start" {
		newGame := NewTicTacToeGame()
		// Optionally set the first player (authorID) as Player X or wait for the first move
		cmd.ActiveGames[m.ChannelID] = newGame
		s.ChannelMessageSend(m.ChannelID, "New Tic Tac Toe game started!")
		return
	}

	// Below this point, handle making a move in an existing game
	game, exists := cmd.ActiveGames[m.ChannelID]
	if !exists {
		s.ChannelMessageSend(m.ChannelID, "No active game. Use `-tictactoe start` to begin a new game.")
		return
	}

	// Ensure there are enough arguments to parse a move
	if len(args) != 3 {
		s.ChannelMessageSend(m.ChannelID, "Invalid move. Please use `-tictactoe x y` with x, y in [0, 2].")
		return
	}

	x, err := strconv.Atoi(args[1])
	y, err2 := strconv.Atoi(args[2])
	if err != nil || err2 != nil || x < 0 || y < 0 || x > 2 || y > 2 {
		s.ChannelMessageSend(m.ChannelID, "Invalid move. Please use `-tictactoe x y` with x, y in [0, 2].")
		return
	}

	err = game.PlayMove(x, y, m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	// Check for a win or continue game
	if game.Winner != Empty {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Player %s wins!\n%s", game.Winner, game.BoardString()))
		delete(cmd.ActiveGames, m.ChannelID)
		return
	}

	// Bot makes a move
	cmd.botMove(game, s, m)

}

func (cmd *TicTacToeCommand) botMove(game *TicTacToeGame, s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(time.Now().UnixNano())
	var x, y int
	for {
		x = rand.Intn(3)
		y = rand.Intn(3)
		if game.Board[x][y] == Empty {
			_ = game.PlayMove(x, y, m.Author.ID)
			break
		}
	}

	// Check for win or tie after bot move
	if game.Winner != Empty {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Bot wins!\n%s", game.BoardString()))
	} else if game.isBoardFull() {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("It's a tie!\n%s", game.BoardString()))
	} else {
		s.ChannelMessageSend(m.ChannelID, "Your move:\n"+game.BoardString())
	}
}
