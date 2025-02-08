package game

import (
	"testing"
)

// TODO use assert and check error messages

func TestPlayer1Wins(t *testing.T) {
	// Create a new board
	board := NewBoard()

	// Player 1 places 3 pieces horizontally in the top row
	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	board.MustMakeMove(NewMove(Player1, 0, 1, Medium))
	board.MustMakeMove(NewMove(Player1, 0, 2, Large))

	// Check for a winner
	winner := board.CheckWin()

	// Test assertion: Player 1 should win
	if winner != Player1 {
		t.Errorf("Expected Player 1 to win, but got %v", winner)
	}
}

func TestOutOfBoundsMove(t *testing.T) {
	board := NewBoard()
	err := board.MakeMove(NewMove(Player1, 3, 0, Small))
	if err == nil {
		t.Errorf("Move should be out of bounds")
	}

	err = board.MakeMove(NewMoveExisting(Player1, 0, 3, Small, 0, 0))
	if err == nil {
		t.Errorf("Move should be out of bounds")
	}
}

func TestPieceNotLarger(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	err := board.MakeMove(NewMove(Player1, 0, 0, Small))
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}

	board.MustMakeMove(Move{Piece: Piece{Owner: Player1, Size: Medium}, To: Position{0, 0}})
	err = board.MakeMove(NewMove(Player1, 0, 0, Small))
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}
}

func TestPieceNotAvailable(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	board.MustMakeMove(NewMove(Player1, 0, 1, Small))
	err := board.MakeMove(NewMove(Player1, 0, 2, Small))
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}
}

func TestMovingPiece(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	board.MustMakeMove(NewMove(Player1, 0, 0, Medium))
	board.MustMakeMove(NewMove(Player1, 2, 2, Small))
	board.MustMakeMove(NewMoveExisting(Player1, 0, 0, Medium, 2, 2))
}

func TestMovingNonExistingPiece(t *testing.T) {
	board := NewBoard()
	err := board.MakeMove(NewMoveExisting(Player1, 0, 0, Small, 1, 1))
	if err == nil {
		t.Errorf("Moving non-existing piece must be illegal")
	}

	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	err = board.MakeMove(NewMoveExisting(Player1, 0, 0, Medium, 1, 1))
	if err == nil {
		t.Errorf("Moving non-existing piece must be illegal")
	}
}

func TestMovingPieceToSameLocation(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	err := board.MakeMove(NewMoveExisting(Player1, 0, 0, Small, 0, 0))
	if err == nil {
		t.Errorf("Moving piece to same location must be illegal")
	}
}

func TestMovingPieceOfOtherPlayer(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player2, 0, 0, Small))

	err := board.MakeMove(NewMoveExisting(Player1, 0, 0, Small, 1, 1))
	if err == nil {
		t.Errorf("Moving piece of other player must be illegal")
	}
}

func TestMovingPieceWouldCauseOtherPlayerToWin(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	board.MustMakeMove(NewMove(Player1, 0, 1, Small))
	board.MustMakeMove(NewMove(Player2, 0, 1, Medium))
	board.MustMakeMove(NewMove(Player1, 0, 2, Large))
	if board.CheckWin() != None {
		t.Errorf("No player should have won yet")
	}

	err := board.MakeMove(NewMoveExisting(Player2, 0, 1, Medium, 2, 2))
	if err == nil {
		t.Errorf("Moving piece that would cause the other player to win is illegal")
	}
}

func TestGetPossibleMoves(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	board.MustMakeMove(NewMove(Player1, 0, 1, Small))
	board.MustMakeMove(NewMove(Player2, 0, 0, Medium))
	board.MustMakeMove(NewMove(Player2, 0, 1, Medium))
	board.MustMakeMove(NewMove(Player1, 0, 0, Large))
	board.MustMakeMove(NewMove(Player1, 0, 1, Large))

	board.MustMakeMove(NewMove(Player2, 1, 0, Small))
	board.MustMakeMove(NewMove(Player1, 1, 0, Medium))
	board.MustMakeMove(NewMove(Player2, 1, 0, Large))

	board.MustMakeMove(NewMove(Player2, 2, 0, Small))

	possibleMoves := board.GetPossibleMoves(Player1)
	if len(possibleMoves) != 12 {
		t.Errorf("GetPossibleMoves() should have returned 12 possible moves")
	}

	board.MustMakeMove(NewMove(Player1, 2, 0, Medium))

	possibleMoves = board.GetPossibleMoves(Player2)
	if len(possibleMoves) != 6 {
		t.Errorf("GetPossibleMoves() should have returned 6 possible moves")
	}

}
