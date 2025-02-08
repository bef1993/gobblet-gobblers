package ai

import "gibhub.com/bef1993/gobblet-gobblers/game"

type Evaluation int

const (
	Player1Win Evaluation = 1
	Player2Win Evaluation = -1
	NoWin      Evaluation = 0
)

func Evaluate(b *game.Board) Evaluation {
	switch b.CheckWin() {
	case game.Player1:
		return Player1Win
	case game.Player2:
		return Player2Win
	case game.None:
		return NoWin
	default:
		panic("game state could not be evaluated")
	}
}
