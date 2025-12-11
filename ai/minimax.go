package ai

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"gibhub.com/bef1993/gobblet-gobblers/game"
)

// Minimax is the interface for the minimax algorithm
type Minimax interface {
	SolvePosition(board *game.Board, maxDepth int) (winner game.Player)
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

func (m *minimax) SolvePosition(board *game.Board, maxDepth int) (winner game.Player) {
	evaluation, _ := m.minimax(board, maxDepth, math.MinInt, math.MaxInt, isMaximizingPlayer(board.ActivePlayer))
	if evaluation == NoWin {
		return game.None
	} else if evaluation >= Player1Win {
		return game.Player1
	} else {
		return game.Player2
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

	for _, possibleMove := range shuffleMoves(board.GetPossibleMoves()) {
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

func shuffleMoves(moves []game.Move) []game.Move {
	// TODO sort moves by heuristic instead of random shuffle
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuffledMoves := make([]game.Move, len(moves))
	perm := r.Perm(len(moves))
	for i, randIndex := range perm {
		shuffledMoves[i] = moves[randIndex]
	}
	return shuffledMoves
}
