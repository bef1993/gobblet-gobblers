package game

import "testing"

func TestZobristHash(t *testing.T) {
	board := NewBoard()
	hash1 := board.Hash
	if hash1 != GetPlayerZobristValue(Player1) {
		t.Error("hash must equal player 1 hash")
	}
	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	hash2 := board.Hash
	board.MustUndoMove(NewMove(Player1, 0, 0, Small))
	if board.Hash != hash1 {
		t.Error("hash must equal after UndoMove")
	}
	board.MustMakeMove(NewMove(Player1, 0, 0, Small))
	if board.Hash != hash2 {
		t.Error("hash must equal after redo move")
	}

	if hash1 == hash2 {
		t.Error("hash must be different for different game states")
	}
}

func TestZobristHashMovePiece(t *testing.T) {
	board := NewBoard()
	board.MustMakeMove(NewMove(Player1, 0, 0, Medium))
	board.MustMakeMove(NewMove(Player2, 1, 1, Small))
	hash1 := board.Hash
	board.MustMakeMove(NewMoveExisting(0, 0, 1, 1))
	hash2 := board.Hash
	board.MustUndoMove(NewMoveExisting(0, 0, 1, 1))
	if board.Hash != hash1 {
		t.Error("hash must equal after UndoMove")
	}
	board.MustMakeMove(NewMoveExisting(0, 0, 1, 1))
	if board.Hash != hash2 {
		t.Error("hash must equal after redo move")
	}

	if hash1 == hash2 {
		t.Error("hash must be different for different game states")
	}
}
