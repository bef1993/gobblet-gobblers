package game

import (
	"testing"
)

func TestPlayer1Wins(t *testing.T) {
	// Create a new board
	board := NewBoard()

	// Player 1 places 3 pieces horizontally in the top row
	board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{0, 1}})
	board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Large}, To: Position{0, 2}})

	// Check for a winner
	winner := board.CheckWin()

	// Test assertion: Player 1 should win
	if winner != Player1 {
		t.Errorf("Expected Player 1 to win, but got %v", winner)
	}
}

func TestOutOfBoundsMove(t *testing.T) {
	board := NewBoard()
	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{3, 0}})
	if err == nil {
		t.Errorf("Move should be out of bounds")
	}

	err = board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}, From: &Position{0, 3}})
	if err == nil {
		t.Errorf("Move should be out of bounds")
	}
}

func TestSimpleInvalidMoves(t *testing.T) {
	board := NewBoard()
	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	if err != nil {
		t.Errorf("Valid move must not be illegal")
	}

	err = board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}

	err = board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{0, 0}})
	if err != nil {
		t.Errorf("Valid move must not be illegal")
	}

	err = board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}
}

func TestMovingNonExistingPiece(t *testing.T) {
	board := NewBoard()
	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{1, 1}, From: &Position{0, 0}})
	if err == nil {
		t.Errorf("Moving non-existing piece must be illegal")
	}
}

func TestMovingPieceToSameLocation(t *testing.T) {
	board := NewBoard()
	board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}, From: &Position{0, 0}})
	if err == nil {
		t.Errorf("Moving piece to same location must be illegal")
	}
}

func TestMovingPieceOfOtherPlayer(t *testing.T) {
	board := NewBoard()
	board.MakeMove(Move{Piece: Piece{Owner: Player2, Size: Small}, To: Position{0, 0}})

	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{1, 1}, From: &Position{0, 0}})
	if err == nil {
		t.Errorf("Moving piece of other player must be illegal")
	}
}

func TestMovingPieceWouldCauseOtherPlayerToWin(t *testing.T) {
	board := NewBoard()
	board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 1}})
	board.MakeMove(Move{Piece: Piece{Owner: Player2, Size: Medium}, To: Position{0, 1}})
	board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Large}, To: Position{0, 2}})
	if board.CheckWin() != None {
		t.Errorf("No player should have won yet")
	}

	err := board.MakeMove(Move{Piece: Piece{Owner: Player2, Size: Medium}, To: Position{2, 2}, From: &Position{0, 1}})
	if err == nil {
		t.Errorf("Moving piece that would cause the other player to win is illegal")
	}
}
