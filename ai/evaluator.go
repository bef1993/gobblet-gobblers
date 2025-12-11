package ai

import "gibhub.com/bef1993/gobblet-gobblers/game"

const (
	Player1Win int = 1000
	Player2Win int = -1000
	NoWin      int = 0
)

type Evaluator interface {
	Evaluate(board *game.Board, depthRemaining int) (evaluation int)
}

type evaluator struct{}

func NewEvaluator() Evaluator {
	return &evaluator{}
}

func (e *evaluator) Evaluate(b *game.Board, depthRemaining int) int {
	switch b.CheckWin() {
	case game.Player1:
		return Player1Win + depthRemaining // prefer faster wins
	case game.Player2:
		return Player2Win - depthRemaining // prefer delaying losses
	case game.None:
		return NoWin
	default:
		panic("game state could not be evaluated")
	}
}
