package ai

import (
	"testing"

	"gibhub.com/bef1993/gobblet-gobblers/game"
	"github.com/stretchr/testify/assert"
)

func TestNoWin(t *testing.T) {
	board := game.NewBoard()
	minimax := NewMinimax()
	winner := minimax.CalculateWinner(board, 5)
	assert.Equal(t, game.None, winner, "game must not be solved with maxDepth 5")
}

func TestPlayer1Win(t *testing.T) {
	board := game.NewBoard()
	board.MustMakeMove(game.NewMove(game.Player1, board.Get(1, 1), game.Small))
	board.MustMakeMove(game.NewMove(game.Player2, board.Get(1, 0), game.Medium))
	board.MustMakeMove(game.NewMove(game.Player1, board.Get(1, 1), game.Large))
	minimax := NewMinimax()
	winner := minimax.CalculateWinner(board, 8)

	assert.Equal(t, game.Player1, winner, "game must be solved with maxDepth 8 and previous moves")
}

func TestPlayer2Win(t *testing.T) {
	board := game.NewBoard()
	board.MustMakeMove(game.NewMove(game.Player1, board.Get(1, 0), game.Medium))
	board.MustMakeMove(game.NewMove(game.Player2, board.Get(1, 1), game.Small))
	board.MustMakeMove(game.NewMove(game.Player1, board.Get(0, 1), game.Large))
	board.MustMakeMove(game.NewMove(game.Player2, board.Get(0, 0), game.Small))
	board.MustMakeMove(game.NewMove(game.Player1, board.Get(1, 0), game.Large))
	minimax := NewMinimax()
	winner := minimax.CalculateWinner(board, 1)

	assert.Equal(t, game.Player2, winner, "winner must be Player 1")
}

func TestFullSolveGame(t *testing.T) {
	board := game.NewBoard()
	minimax := NewMinimax()
	winner := minimax.CalculateWinner(board, 9)

	assert.Equal(t, game.Player1, winner, "winner must be Player 1")
}
