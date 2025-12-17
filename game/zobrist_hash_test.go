package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZobristHash(t *testing.T) {
	board := NewBoard()
	hash1 := board.Hash
	assert.Equal(t, GetPlayerZobristValue(Player1), hash1, "hash must be equal to player 1 hash")

	move := NewMove(Player1, board.Get(0, 0), Small)
	board.MustMakeMove(move)
	hash2 := board.Hash
	board.MustUndoMove(move)
	assert.Equal(t, hash1, board.Hash, "hash must be equal after UndoMove")

	board.MustMakeMove(move)
	assert.Equal(t, hash2, board.Hash, "hash must be equal after redo move")

	assert.NotEqual(t, hash1, hash2, "hash must be different for different game states")
}

func TestZobristHashMovePiece(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, board.Get(0, 0), Medium))
	board.MustMakeMove(NewMove(Player2, board.Get(1, 1), Small))
	hash1 := board.Hash
	move := NewMoveExisting(board.Get(0, 0), board.Get(1, 1))
	board.MustMakeMove(move)
	hash2 := board.Hash
	board.MustUndoMove(move)
	assert.Equal(t, hash1, board.Hash, "hash must be equal after UndoMove")

	board.MustMakeMove(move)
	assert.Equal(t, hash2, board.Hash, "hash must be equal after redo move")

	assert.NotEqual(t, hash1, hash2, "hash must be different for different game states")
}
