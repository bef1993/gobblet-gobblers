package game

import (
	"testing"
)

// TODO use assert and check error messages

func TestPlayer1Wins(t *testing.T) {
	// Create a new board
	board := NewBoard()

	// Player 1 places 3 pieces horizontally in the top row
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{0, 1}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Large}, To: Position{0, 2}})

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

func TestPieceNotLarger(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}

	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{0, 0}})
	err = board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}
}

func TestPieceNotAvailable(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 1}})
	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 2}})
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}
}

func TestMovingPiece(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{0, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{2, 2}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{2, 2}, From: &Position{0, 0}})
}

func TestMovingNonExistingPiece(t *testing.T) {
	board := NewBoard()
	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{1, 1}, From: &Position{0, 0}})
	if err == nil {
		t.Errorf("Moving non-existing piece must be illegal")
	}

	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	err = board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{1, 1}, From: &Position{0, 0}})
	if err == nil {
		t.Errorf("Moving non-existing piece must be illegal")
	}
}

func TestMovingPieceToSameLocation(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}, From: &Position{0, 0}})
	if err == nil {
		t.Errorf("Moving piece to same location must be illegal")
	}
}

func TestMovingPieceOfOtherPlayer(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(Move{Piece: Piece{Owner: Player2, Size: Small}, To: Position{0, 0}})

	err := board.MakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{1, 1}, From: &Position{0, 0}})
	if err == nil {
		t.Errorf("Moving piece of other player must be illegal")
	}
}

func TestMovingPieceWouldCauseOtherPlayerToWin(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 1}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player2, Size: Medium}, To: Position{0, 1}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Large}, To: Position{0, 2}})
	if board.CheckWin() != None {
		t.Errorf("No player should have won yet")
	}

	err := board.MakeMove(Move{Piece: Piece{Owner: Player2, Size: Medium}, To: Position{2, 2}, From: &Position{0, 1}})
	if err == nil {
		t.Errorf("Moving piece that would cause the other player to win is illegal")
	}
}

func TestGetPossibleMoves(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Small}, To: Position{0, 1}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player2, Size: Medium}, To: Position{0, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player2, Size: Medium}, To: Position{0, 1}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Large}, To: Position{0, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Large}, To: Position{0, 1}})

	board.MustMakeMove(Move{Piece: Piece{Owner: Player2, Size: Small}, To: Position{1, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{1, 0}})
	board.MustMakeMove(Move{Piece: Piece{Owner: Player2, Size: Large}, To: Position{1, 0}})

	board.MustMakeMove(Move{Piece: Piece{Owner: Player2, Size: Small}, To: Position{2, 0}})

	possibleMoves := board.GetPossibleMoves(Player1)
	if len(possibleMoves) != 12 {
		t.Errorf("GetPossibleMoves() should have returned 12 possible moves")
	}

	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{2, 0}})

	possibleMoves = board.GetPossibleMoves(Player2)
	if len(possibleMoves) != 6 {
		t.Errorf("GetPossibleMoves() should have returned 6 possible moves")
	}

}
