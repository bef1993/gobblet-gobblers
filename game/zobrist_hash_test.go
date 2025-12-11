package game

import "testing"

func TestZobristHash(t *testing.T) {
	board := NewBoard()
	hash1 := board.Hash
	if hash1 != GetPlayerZobristValue(Player1) {
		t.Error("hash must equal player 1 hash")
	}
	move := NewMove(Player1, board.Get(0, 0), Small)
	board.MustMakeMove(move)
	hash2 := board.Hash
	board.MustUndoMove(move)
	if board.Hash != hash1 {
		t.Error("hash must equal after UndoMove")
	}
	board.MustMakeMove(move)
	if board.Hash != hash2 {
		t.Error("hash must equal after redo move")
	}

	if hash1 == hash2 {
		t.Error("hash must be different for different game states")
	}
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
	if board.Hash != hash1 {
		t.Error("hash must equal after UndoMove")
	}
	board.MustMakeMove(move)
	if board.Hash != hash2 {
		t.Error("hash must equal after redo move")
	}

	if hash1 == hash2 {
		t.Error("hash must be different for different game states")
	}
}
