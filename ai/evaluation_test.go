package ai

import (
	"gibhub.com/bef1993/gobblet-gobblers/game"
	"testing"
)

func TestEvaluationPlayer1Win(t *testing.T) {
	board := game.NewBoard()
	board.MustMakeMove(game.Move{Piece: game.Piece{Owner: game.Player1, Size: game.Small}, To: game.Position{Row: 0, Col: 0}})
	board.MustMakeMove(game.Move{Piece: game.Piece{Owner: game.Player1, Size: game.Medium}, To: game.Position{Row: 1, Col: 0}})
	board.MustMakeMove(game.Move{Piece: game.Piece{Owner: game.Player1, Size: game.Small}, To: game.Position{Row: 2, Col: 0}})
	if Evaluate(board) != Player1Win {
		t.Errorf("Game must be evaluated to Player1Win")
	}
}

func TestEvaluationPlayer2Win(t *testing.T) {
	board := game.NewBoard()
	board.MustMakeMove(game.Move{Piece: game.Piece{Owner: game.Player2, Size: game.Small}, To: game.Position{Row: 0, Col: 0}})
	board.MustMakeMove(game.Move{Piece: game.Piece{Owner: game.Player2, Size: game.Medium}, To: game.Position{Row: 1, Col: 1}})
	board.MustMakeMove(game.Move{Piece: game.Piece{Owner: game.Player2, Size: game.Small}, To: game.Position{Row: 2, Col: 2}})
	if Evaluate(board) != Player2Win {
		t.Errorf("Game must be evaluated to Player2Win")
	}
}

func TestEvaluationNoWin(t *testing.T) {
	board := game.NewBoard()
	board.MustMakeMove(game.Move{Piece: game.Piece{Owner: game.Player1, Size: game.Small}, To: game.Position{Row: 0, Col: 0}})
	board.MustMakeMove(game.Move{Piece: game.Piece{Owner: game.Player2, Size: game.Small}, To: game.Position{Row: 1, Col: 0}})
	board.MustMakeMove(game.Move{Piece: game.Piece{Owner: game.Player1, Size: game.Small}, To: game.Position{Row: 2, Col: 0}})
	if Evaluate(board) != NoWin {
		t.Errorf("Game must be evaluated to NoWin")
	}
}
