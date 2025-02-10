package ai

import (
	"gibhub.com/bef1993/gobblet-gobblers/game"
	"math"
	"math/rand"
	"time"
)

const (
	Player1Win int = 100
	Player2Win int = -100
	NoWin      int = 0
)

var MaxDepth = 7

func SolvePosition(board *game.Board) (winner game.Player) {
	evaluation, _ := minimax(board, MaxDepth, math.MinInt, math.MaxInt, isMaximizingPlayer(board.ActivePlayer))
	switch evaluation {
	case Player1Win:
		return game.Player1
	case Player2Win:
		return game.Player2
	default:
		return game.None
	}
}

func GetBestMove(board *game.Board) game.Move {
	_, bestMove := minimax(board, MaxDepth, Player2Win, Player1Win, isMaximizingPlayer(board.ActivePlayer))
	return bestMove
}

func minimax(board *game.Board, depth, alpha, beta int, isMaximizingPlayer bool) (evaluation int, bestMove game.Move) {
	hash := Hash(board)

	// Check the Transposition Table first
	if found, evaluation, bestMove := lookupHash(hash, depth, alpha, beta); found {
		return evaluation, bestMove
	}

	if depth == 0 || board.CheckWin() != game.None {
		evaluation := Evaluate(board)
		storeHash(hash, evaluation, depth, Exact)
		return evaluation, game.Move{}
	}

	maxEval := math.MinInt
	minEval := math.MaxInt

	for _, possibleMove := range shuffleMoves(board.GetPossibleMoves()) {
		board.MustMakeMove(possibleMove)
		eval, _ := minimax(board, depth-1, alpha, beta, !isMaximizingPlayer)
		board.MustUndoMove(possibleMove)

		if isMaximizingPlayer {
			alpha = max(alpha, eval)
			if eval > maxEval {
				maxEval = eval
				bestMove = possibleMove
			}
		} else {
			beta = min(beta, eval)
			if eval < minEval {
				minEval = eval
				bestMove = possibleMove
			}
		}
	}
	// TODO prefer moves that win harder / lose slower
	if isMaximizingPlayer {
		storeHash(hash, maxEval, depth, LowerBound)
		return maxEval, bestMove
	} else {
		storeHash(hash, minEval, depth, UpperBound)
		return minEval, bestMove
	}
}

func Evaluate(b *game.Board) int {
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
