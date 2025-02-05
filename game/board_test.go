package game

import (
	"testing"
)

func TestPlayer1Wins(t *testing.T) {
	// Create a new board
	board := NewBoard()

	// Player 1 places 3 pieces horizontally in the top row
	board.PlacePiece(0, 0, Piece{Owner: Player1, Size: Small})
	board.PlacePiece(0, 1, Piece{Owner: Player1, Size: Medium})
	board.PlacePiece(0, 2, Piece{Owner: Player1, Size: Large})

	// Check for a winner
	winner := board.CheckWin()

	// Test assertion: Player 1 should win
	if winner != Player1 {
		t.Errorf("Expected Player 1 to win, but got %v", winner)
	}
}

func TestIllegalMove(t *testing.T) {
	board := NewBoard()
	board.PlacePiece(0, 0, Piece{Owner: Player1, Size: Small})
	err := board.PlacePiece(0, 0, Piece{Owner: Player1, Size: Small})
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}
}
