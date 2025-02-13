package ai

import (
	"fmt"
	"gibhub.com/bef1993/gobblet-gobblers/game"
	"math"
	"math/rand"
	"time"
)

const (
	Player1Win int = 1000
	Player2Win int = -1000
	NoWin      int = 0
)

type EntryType int

const (
	Exact EntryType = iota
	LowerBound
	UpperBound
)

type TTEntry struct {
	Evaluation int
	Depth      int
	BestMove   game.Move
	Type       EntryType
}

var transpositionTable = make(map[uint64]TTEntry)

func SolvePosition(board *game.Board, maxDepth int) (winner game.Player) {
	evaluation, _ := minimax(board, maxDepth, math.MinInt, math.MaxInt, isMaximizingPlayer(board.ActivePlayer))
	if evaluation == NoWin {
		return game.None
	} else if evaluation >= Player1Win {
		return game.Player1
	} else {
		return game.Player2
	}
}

func GetBestMove(board *game.Board, maxDepth int) game.Move {
	eval, bestMove := minimax(board, maxDepth, math.MinInt, math.MaxInt, isMaximizingPlayer(board.ActivePlayer))
	fmt.Printf("Evaluation: %v\n", eval)
	return bestMove
}

func minimax(board *game.Board, depth, alpha, beta int, isMaximizingPlayer bool) (evaluation int, bestMove game.Move) {

	// Check the Transposition Table first
	if found, evaluation, storedMove := lookupHash(board.Hash, depth, alpha, beta); found {
		if valid, _ := board.IsValidMove(storedMove); valid {
			return evaluation, storedMove
		}
	}

	if depth == 0 || board.CheckWin() != game.None {
		evaluation := Evaluate(board, depth)
		storeHash(board.Hash, evaluation, depth, Exact, game.Move{})
		return evaluation, game.Move{}
	}

	maxEval := math.MinInt
	minEval := math.MaxInt

	for _, possibleMove := range shuffleMoves(board.GetPossibleMoves()) {
		board.MustMakeMove(possibleMove)
		eval, _ := minimax(board, depth-1, alpha, beta, !isMaximizingPlayer)
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
		storeHash(board.Hash, maxEval, depth, LowerBound, bestMove)
		return maxEval, bestMove
	} else {
		storeHash(board.Hash, minEval, depth, UpperBound, bestMove)
		return minEval, bestMove
	}
}

func Evaluate(b *game.Board, depthRemaining int) int {
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

func lookupHash(hash uint64, depth, alpha, beta int) (found bool, evaluation int, bestMove game.Move) {
	entry, exists := transpositionTable[hash]
	if !exists || entry.Depth < depth {
		return false, NoWin, game.Move{} // Not found or outdated
	}

	// Use stored value if it helps pruning
	if entry.Type == Exact {
		return true, entry.Evaluation, entry.BestMove
	} else if entry.Type == LowerBound && entry.Evaluation >= beta {
		return true, entry.Evaluation, entry.BestMove
	} else if entry.Type == UpperBound && entry.Evaluation <= alpha {
		return true, entry.Evaluation, entry.BestMove
	}

	return false, NoWin, game.Move{}
}

func storeHash(hash uint64, evaluation, depth int, entryType EntryType, bestMove game.Move) {
	existing, exists := transpositionTable[hash]

	// Only replace if the new depth is greater, or it's a new entry
	if !exists || depth > existing.Depth {
		transpositionTable[hash] = TTEntry{Evaluation: evaluation, Depth: depth, Type: entryType, BestMove: bestMove}
	}
}
