package ai

import (
	"gibhub.com/bef1993/gobblet-gobblers/game"
	"testing"
)

// TODO add tests

func TestNoWin(t *testing.T) {
	board := game.NewBoard()
	winner := SolvePosition(board, 5)
	if winner != game.None {
		t.Errorf("game is not winnable with maxDepth 5")
	}
}

func TestPlayer1Win(t *testing.T) {
	board := game.NewBoard()
	board.MustMakeMove(game.NewMove(game.Player1, 1, 1, game.Small))
	board.MustMakeMove(game.NewMove(game.Player2, 1, 0, game.Medium))
	board.MustMakeMove(game.NewMove(game.Player1, 1, 1, game.Large))
	winner := SolvePosition(board, 8)
	if winner != game.Player1 {
		t.Errorf("game is winnable with maxDepth 8")
	}
}

func TestPlayer2Win(t *testing.T) {
	board := game.NewBoard()
	board.MustMakeMove(game.NewMove(game.Player1, 1, 0, game.Medium))
	board.MustMakeMove(game.NewMove(game.Player2, 1, 1, game.Small))
	board.MustMakeMove(game.NewMove(game.Player1, 0, 1, game.Large))
	board.MustMakeMove(game.NewMove(game.Player2, 0, 0, game.Small))
	board.MustMakeMove(game.NewMove(game.Player1, 1, 0, game.Large))
	winner := SolvePosition(board, 1)
	if winner != game.Player2 {
		t.Errorf("winner must be Player 1")
	}
}
