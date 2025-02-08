package ai

import (
	"math"

	"gibhub.com/bef1993/gobblet-gobblers/game"
)

type AIPlayer struct {
	Player game.Player
}

func NewAIPlayer(player game.Player) *AIPlayer {
	return &AIPlayer{Player: player}
}

func (ai *AIPlayer) BestMove(board *game.Board) game.Move {
	bestScore := math.Inf(-1)
	var bestMove game.Move

	moveStack := []game.Move{}

	possibleMoves := board.GetPossibleMoves(ai.Player)
	for _, move := range possibleMoves {
		board.MustMakeMove(move)
		moveStack = append(moveStack, move)
		score := ai.minimax(board, 0, math.Inf(-1), math.Inf(1), false)
		moveStack = moveStack[:len(moveStack)-1]

		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove
}

func (ai *AIPlayer) minimax(board *game.Board, depth int, alpha, beta float64, maximizingPlayer bool) float64 {
	winner := board.CheckWin()
	if winner != game.None {
		return ai.evaluateWinner(winner, depth)
	}

	moveStack := []game.Move{}

	if maximizingPlayer {
		maxEval := math.Inf(-1)
		for _, move := range board.GetPossibleMoves(ai.Player) {
			board.MustMakeMove(move)
			moveStack = append(moveStack, move)
			eval := ai.minimax(board, depth+1, alpha, beta, false)
			moveStack = moveStack[:len(moveStack)-1]
			maxEval = math.Max(maxEval, eval)
			alpha = math.Max(alpha, eval)
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		minEval := math.Inf(1)
		opponent := game.Player1
		if ai.Player == game.Player1 {
			opponent = game.Player2
		}
		for _, move := range board.GetPossibleMoves(opponent) {
			board.MustMakeMove(move)
			moveStack = append(moveStack, move)
			eval := ai.minimax(board, depth+1, alpha, beta, true)
			moveStack = moveStack[:len(moveStack)-1]
			minEval = math.Min(minEval, eval)
			beta = math.Min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}

func (ai *AIPlayer) evaluateWinner(winner game.Player, depth int) float64 {
	if winner == ai.Player {
		return 10 - float64(depth)
	} else if winner == game.None {
		return 0
	}
	return float64(depth) - 10
}
