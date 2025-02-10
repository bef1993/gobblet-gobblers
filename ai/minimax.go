package ai

import (
	"gibhub.com/bef1993/gobblet-gobblers/game"
	"math"
)

const (
	Player1Win float64 = 10
	Player2Win float64 = -10
	NoWin      float64 = 0
	maxDepth   float64 = 5
)

func GetBestMove(board *game.Board, activePlayer game.Player) game.Move {
	_, bestMove := minimax(board, maxDepth, math.Inf(-1), math.Inf(1), isMaximizingPlayer(activePlayer), activePlayer)
	return bestMove
}

func minimax(board *game.Board, depth, alpha, beta float64, isMaximizingPlayer bool, activePlayer game.Player) (evaluation float64, bestMove game.Move) {
	if depth == 0 || board.CheckWin() != game.None {
		return Evaluate(board), game.Move{}
	}

	maxEval := math.Inf(-1)
	minEval := math.Inf(1)

	for _, possibleMove := range board.GetPossibleMoves(activePlayer) {
		// TODO sort moves by heuristic
		board.MustMakeMove(possibleMove)
		// TODO use transposition table to avoid recomputing identical positions (zobrist hashing)
		eval, _ := minimax(board, depth-1, alpha, beta, !isMaximizingPlayer, activePlayer.Opponent())
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

	if isMaximizingPlayer {
		evaluation = maxEval
	} else {
		evaluation = minEval
	}
	return evaluation, bestMove
}

func Evaluate(b *game.Board) float64 {
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
