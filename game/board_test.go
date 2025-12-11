package game

import (
	"testing"
)

// TODO use assert and check error messages

func TestPlayer1Win(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Large))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 0), Small))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 0), Medium))
	board.MustMakeMove(NewMove(Player2, board.Get(2, 2), Small))
	board.MustMakeMove(NewMove(Player1, board.Get(2, 0), Small))

	if board.CheckWin() != Player1 {
		t.Errorf("Game must be won by player 1")
	}

	if len(board.GetPossibleMoves()) > 0 {
		t.Errorf("Game must have no possible moves because it is already won")
	}
}

func TestPlayer2Win(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 1), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(2, 2), Medium))
	board.MustMakeMove(NewMove(Player2, board.Get(2, 2), Large))

	if board.CheckWin() != Player2 {
		t.Errorf("Game must be won by player 2")
	}
}

func TestNoWin(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Small))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 1), Medium))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Large))
	board.MustMakeMove(NewMove(Player1, board.Get(2, 2), Small))
	if board.CheckWin() != None {
		t.Errorf("Game must not be won yet")
	}
}

func TestPlaceOpponentPiece(t *testing.T) {
	board := NewBoard()
	err := board.MakeMove(NewMove(Player2, board.Get(0, 0), Small))
	if err == nil {
		t.Errorf("placing piece of opponent must be illegal")
	}
}

func TestPieceNotLarger(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	err := board.MakeMove(NewMove(Player2, board.Get(0, 0), Small))
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}

	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	err = board.MakeMove(NewMove(Player1, board.Get(0, 0), Medium))
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}
}

func TestPieceNotAvailable(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 1), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Medium))
	err := board.MakeMove(NewMove(Player1, board.Get(2, 2), Small))
	if err == nil {
		t.Errorf("Expected move to be illegal")
	}
}

func TestMovingPiece(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Large))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Medium))
	board.MustMakeMove(NewMoveExisting(board.Get(0, 0), board.Get(1, 1)))
}

func TestMovingNonExistingPiece(t *testing.T) {
	board := NewBoard()
	err := board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(1, 1)))
	if err == nil {
		t.Errorf("Moving non-existing piece must be illegal")
	}

	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	err = board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(1, 1)))
	if err == nil {
		t.Errorf("Moving non-existing piece must be illegal")
	}
}

func TestMovingPieceToSameLocation(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Small))
	err := board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(0, 0)))
	if err == nil {
		t.Errorf("Moving piece to same location must be illegal")
	}
}

func TestMovingPieceOfOtherPlayer(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))

	err := board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(1, 1)))
	if err == nil {
		t.Errorf("Moving piece of other player must be illegal")
	}
}

func TestMovingPieceWouldCauseOtherPlayerToWin(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 1), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(2, 2), Small))
	board.MustMakeMove(NewMove(Player1, board.Get(2, 2), Medium))
	err := board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(0, 2)))
	if err == nil {
		t.Errorf("Moving piece that would cause the other player to win is illegal")
	}
}

func TestGetPossibleMoves(t *testing.T) {
	board := NewBoard()
	if len(board.GetPossibleMoves()) != 9*3 {
		t.Errorf("GetPossibleMoves should have 9*3 possible moves")
	}
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	if len(board.GetPossibleMoves()) != (9*3)-1 {
		t.Errorf("GetPossibleMoves should have (9*3)-1 possible moves")
	}
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Small))
	if len(board.GetPossibleMoves()) != (9*3)-2+7 {
		t.Errorf("GetPossibleMoves should have (9*3)-2+7 possible moves")
	}
}

func TestUndoMove(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))

}
