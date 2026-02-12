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

func TestLossDueToNoLegalMoves(t *testing.T) {
	board := NewBoard()

	// Sequence to reach a state where Player 1 has no legal moves.
	// P1 pieces are pinned by P2 threats, and P1 has no pieces in hand.

	// 1. P1: S -> (2,0)
	board.MustMakeMove(NewMove(Player1, board.Get(2, 0), Small))
	// 2. P2: S -> (0,1)
	board.MustMakeMove(NewMove(Player2, board.Get(0, 1), Small))
	// 3. P1: S -> (2,2)
	board.MustMakeMove(NewMove(Player1, board.Get(2, 2), Small))
	// 4. P2: S -> (1,1)
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Small))

	// 5. P1: M -> (0,0)
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Medium))
	// 6. P2: M -> (2,0) [Gobbles P1(S)]
	board.MustMakeMove(NewMove(Player2, board.Get(2, 0), Medium))
	// 7. P1: M -> (0,2)
	board.MustMakeMove(NewMove(Player1, board.Get(0, 2), Medium))
	// 8. P2: M -> (2,2) [Gobbles P1(S)]
	board.MustMakeMove(NewMove(Player2, board.Get(2, 2), Medium))

	// 9. P1: L -> (1,1) [Gobbles P2(S)]
	board.MustMakeMove(NewMove(Player1, board.Get(1, 1), Large))
	// 10. P2: L -> (0,0) [Gobbles P1(M)]
	board.MustMakeMove(NewMove(Player2, board.Get(0, 0), Large))
	// 11. P1: L -> (0,1) [Gobbles P2(S)]
	board.MustMakeMove(NewMove(Player1, board.Get(0, 1), Large))
	// 12. P2: L -> (0,2) [Gobbles P1(M)]
	board.MustMakeMove(NewMove(Player2, board.Get(0, 2), Large))

	// Verify State:
	// P1 has no pieces in hand.
	// P1 has 2 pieces on board: (0,1)L and (1,1)L.
	// (0,1)L covers P2(S). Lifting it exposes Row 0 win for P2 (L-S-L).
	// (1,1)L covers P2(S). Lifting it exposes Diag wins for P2.

	assert.Equal(t, Player1, board.ActivePlayer, "It should be Player 1's turn")
	assert.Equal(t, Player2, board.CheckWin(), "Player 1 has no legal moves, so Player 2 should win")
	assert.False(t, board.HasAnyLegalMove(), "Player 1 should have no legal moves")
	assert.Equal(t, 0, len(board.GetPossibleMoves()), "Player 1 should have 0 possible moves")
}
