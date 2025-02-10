package ai

import (
	"gibhub.com/bef1993/gobblet-gobblers/game"
	"math/rand"
)

var zobristTable [3][3][6]uint64 // 3x3 board, 6 unique pieces
var activePlayerHash [3]uint64   // 1 random value per player
var transpositionTable = make(map[uint64]TTEntry)

type TTEntry struct {
	Evaluation int
	Depth      int
	BestMove   game.Move
	Type       EntryType
}

type EntryType int

const (
	Exact EntryType = iota
	LowerBound
	UpperBound
)

// InitZobrist initializes the hash table
func InitZobrist() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			for piece := 0; piece < 6; piece++ {
				zobristTable[row][col][piece] = rand.Uint64()
			}
		}
	}
	activePlayerHash[1] = rand.Uint64() // Player 1 hash
	activePlayerHash[2] = rand.Uint64() // Player 2 hash
}

// Hash computes the board hash including the player's turn
func Hash(b *game.Board) uint64 {
	var hash uint64 = 0

	// Iterate over all board positions
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			stack := b.Grid[row][col]
			for _, piece := range stack { // Include all pieces in stack
				hash ^= zobristTable[row][col][piece.ID()] // XOR each piece's hash
			}
		}
	}

	// Include the player's turn
	hash ^= activePlayerHash[b.ActivePlayer]

	return hash
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

func storeHash(hash uint64, evaluation, depth int, entryType EntryType) {
	existing, exists := transpositionTable[hash]

	// Only replace if the new depth is greater, or it's a new entry
	if !exists || depth > existing.Depth {
		transpositionTable[hash] = TTEntry{Evaluation: evaluation, Depth: depth, Type: entryType}
	}
}

func init() {
	InitZobrist() // Runs automatically before `main()`
}
