package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer1Win(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Large))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 0), Small))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 0), Medium))
	board.MustMakeMove(NewMove(Player2, board.Get(2, 2), Small))
	board.MustMakeMove(NewMove(Player1, board.Get(2, 0), Small))

	assert.Equal(t, Player1, board.CheckWin(), "Player 1 must win")

	assert.True(t, len(board.GetPossibleMoves()) == 0, "Game must not have possible moves because it is already over")
}

func TestPlayer2Win(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 1), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(2, 2), Medium))
	board.MustMakeMove(NewMove(Player2, board.Get(2, 2), Large))

	assert.Equal(t, Player2, board.CheckWin(), "Player 2 must win")
}

func TestNoWin(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Small))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 1), Medium))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Large))
	board.MustMakeMove(NewMove(Player1, board.Get(2, 2), Small))
	assert.Equal(t, None, board.CheckWin(), "Game must not be won yet")
}

func TestPlaceOpponentPiece(t *testing.T) {
	board := NewBoard()
	assert.Error(t, board.MakeMove(NewMove(Player2, board.Get(0, 0), Small)), "placing opponent piece must be illegal")
}

func TestPieceNotLarger(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	assert.Error(t, board.MakeMove(NewMove(Player2, board.Get(0, 0), Small)), "Expected move to be illegal")

	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	assert.Error(t, board.MakeMove(NewMove(Player1, board.Get(0, 0), Medium)), "Expected move to be illegal")
}

func TestPieceNotAvailable(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 1), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Medium))
	assert.Error(t, board.MakeMove(NewMove(Player1, board.Get(2, 2), Small)), "Expected move to be illegal")
}

func TestMovingPiece(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Large))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Medium))
	board.MustMakeMove(NewMoveExisting(board.Get(0, 0), board.Get(1, 1)))
}

func TestMovingNonExistingPiece(t *testing.T) {
	board := NewBoard()
	assert.Error(t, board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(1, 1))), "\"Moving non-existing piece must be illegal\"")

	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	assert.Error(t, board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(1, 1))), "Moving non-existing piece must be illegal")
}

func TestMovingPieceToSameLocation(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Small))
	assert.Error(t, board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(0, 0))), "Moving piece to same location must be illegal")
}

func TestMovingPieceOfOtherPlayer(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))

	assert.Error(t, board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(1, 1))), "Moving piece of other player must be illegal")
}

func TestMovingPieceWouldCauseOtherPlayerToWin(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Medium))
	board.MustMakeMove(NewMove(Player1, board.Get(1, 1), Small))
	board.MustMakeMove(NewMove(Player2, board.Get(2, 2), Small))
	board.MustMakeMove(NewMove(Player1, board.Get(2, 2), Medium))

	assert.Error(t, board.MakeMove(NewMoveExisting(board.Get(0, 0), board.Get(0, 2))), "Moving piece that would cause the other player to win is illegal")
}

func TestGetPossibleMoves(t *testing.T) {
	board := NewBoard()
	assert.Equal(t, 9*3, len(board.GetPossibleMoves()), "GetPossibleMoves should have 9*3 possible moves")

	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))
	assert.Equal(t, (9*3)-1, len(board.GetPossibleMoves()), "GetPossibleMoves should have (9*3)-1 possible moves")

	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Small))
	assert.Equal(t, (9*3)-2+7, len(board.GetPossibleMoves()), "GetPossibleMoves should have (9*3)-2+7 possible moves")
}

func TestUndoMove(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Small))

}
