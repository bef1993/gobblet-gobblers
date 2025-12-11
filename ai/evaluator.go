package ai

import "gibhub.com/bef1993/gobblet-gobblers/game"

const (
	Player1Win int = 1000
	Player2Win int = -1000
	NoWin      int = 0
)

type Evaluator interface {
	Evaluate(board *game.Board, depthRemaining int) (evaluation int)
	EvaluateMove(board *game.Board, move game.Move) (evaluation int)
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
		// Heuristic evaluation for non-terminal positions
		return e.calculateHeuristicScore(b)
	default:
		panic("game state could not be evaluated")
	}
}

func (e *evaluator) EvaluateMove(b *game.Board, move game.Move) int {
	b.MustMakeMove(move)
	score := e.calculateHeuristicScore(b)
	b.MustUndoMove(move)
	return score
}

func (e *evaluator) calculateHeuristicScore(b *game.Board) int {
	score := 0

	for _, line := range b.Lines {
		score += evaluateLine(line)
	}

	return score
}

func evaluateLine(line game.Line) int {
	p1PieceCount := 0
	p2PieceCount := 0

	for _, pos := range line {
		if pos.TopPiece() != nil {
			if pos.TopPiece().Owner == game.Player1 {
				p1PieceCount++
			} else {
				p2PieceCount++
			}
		}
	}

	// Pieces from both players in the same line
	if p1PieceCount > 0 && p2PieceCount > 0 {
		return 0 // Mixed line, no potential
	}

	// Player 1 has potential
	if p1PieceCount == 2 {
		return 100
	}
	if p1PieceCount == 1 {
		return 10
	}

	// Player 2 has potential
	if p2PieceCount == 2 {
		return -100
	}
	if p2PieceCount == 1 {
		return -10
	}

	return 0 // Empty line
}
