package ai

import (
	"fmt"
	"math"
	"sort"

	"gibhub.com/bef1993/gobblet-gobblers/game"
)

// Minimax is the interface for the minimax algorithm
type Minimax interface {
	CalculateWinner(board *game.Board, maxDepth int) (winner game.Player)
	GetBestMove(board *game.Board, maxDepth int) game.Move
}

// minimax struct holds the state for the minimax algorithm
type minimax struct {
	ttable    TranspositionTable
	evaluator Evaluator
}

// NewMinimax creates a new Minimax instance
func NewMinimax() Minimax {
	return &minimax{
		ttable:    NewTranspositionTable(),
		evaluator: NewEvaluator(),
	}
}

func (m *minimax) CalculateWinner(board *game.Board, maxDepth int) (winner game.Player) {
	evaluation, _ := m.minimax(board, maxDepth, math.MinInt, math.MaxInt, isMaximizingPlayer(board.ActivePlayer))
	if evaluation >= Player1Win {
		return game.Player1
	} else if evaluation <= Player2Win {
		return game.Player2
	} else {
		return game.None
	}
}

func (m *minimax) GetBestMove(board *game.Board, maxDepth int) game.Move {
	eval, bestMove := m.minimax(board, maxDepth, math.MinInt, math.MaxInt, isMaximizingPlayer(board.ActivePlayer))
	fmt.Printf("Evaluation: %v\n", eval)
	return bestMove
}

func (m *minimax) minimax(board *game.Board, depth, alpha, beta int, isMaximizingPlayer bool) (evaluation int, bestMove game.Move) {

	// Check the Transposition Table first
	if found, evaluation, storedMove := m.ttable.LookupHash(board.Hash, depth, alpha, beta); found {
		if valid, _ := board.IsValidMove(storedMove); valid {
			return evaluation, storedMove
		}
	}

	if depth == 0 || board.CheckWin() != game.None {
		evaluation := m.evaluator.Evaluate(board, depth)
		m.ttable.StoreHash(board.Hash, evaluation, depth, ExactBound, game.Move{})
		return evaluation, game.Move{}
	}

	maxEval := math.MinInt
	minEval := math.MaxInt

	sortedMoves := m.sortMoves(board, board.GetPossibleMoves(), isMaximizingPlayer)

	for _, possibleMove := range sortedMoves {
		board.MustMakeMove(possibleMove)
		eval, _ := m.minimax(board, depth-1, alpha, beta, !isMaximizingPlayer)
		board.MustUndoMove(possibleMove)

		if isMaximizingPlayer {
			if eval > maxEval {
				maxEval = eval
				bestMove = possibleMove
			}
			alpha = max(alpha, maxEval)
		} else {
			if eval < minEval {
				minEval = eval
				bestMove = possibleMove
			}
			beta = min(beta, minEval)
		}

		if beta <= alpha {
			break
		}
	}

	if isMaximizingPlayer {
		m.ttable.StoreHash(board.Hash, maxEval, depth, LowerBound, bestMove)
		return maxEval, bestMove
	} else {
		m.ttable.StoreHash(board.Hash, minEval, depth, UpperBound, bestMove)
		return minEval, bestMove
	}
}

func isMaximizingPlayer(player game.Player) bool {
	return player == game.Player1
}

func (m *minimax) sortMoves(board *game.Board, moves []game.Move, isMaximizingPlayer bool) []game.Move {
	moveScores := make(map[game.Move]int)
	for _, move := range moves {
		moveScores[move] = m.evaluator.EvaluateMove(board, move)
	}

	sort.Slice(moves, func(i, j int) bool {
		if isMaximizingPlayer {
			return moveScores[moves[i]] > moveScores[moves[j]]
		}
		return moveScores[moves[i]] < moveScores[moves[j]]
	})

	return moves
}
