package ai

import (
	"gibhub.com/bef1993/gobblet-gobblers/game"
	"testing"
)

// TODO add tests

func TestSolveGame(t *testing.T) {
	board := game.NewBoard()
	winner := SolvePosition(board)
	if winner != game.Player1 {
		t.Errorf("winner must be Player 1")
	}
}
