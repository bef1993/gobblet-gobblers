package ai

import "gibhub.com/bef1993/gobblet-gobblers/game"

type BoundType int

const (
	ExactBound BoundType = iota
	LowerBound
	UpperBound
)

type TTEntry struct {
	Evaluation int
	Depth      int
	BestMove   game.Move
	BoundType  BoundType
}

type TranspositionTable interface {
	LookupHash(hash uint64, depth, alpha, beta int) (found bool, evaluation int, bestMove game.Move)
	StoreHash(hash uint64, evaluation, depth int, entryType BoundType, bestMove game.Move)
}

type transpositionTable struct {
	entries map[uint64]TTEntry
}

func NewTranspositionTable() TranspositionTable {
	return &transpositionTable{
		entries: make(map[uint64]TTEntry),
	}
}

func (t *transpositionTable) LookupHash(hash uint64, depth, alpha, beta int) (found bool, evaluation int, bestMove game.Move) {
	entry, exists := t.entries[hash]
	if !exists || entry.Depth < depth {
		return false, NoWin, game.Move{} // Not found or outdated
	}

	// Use stored value if it helps pruning
	if entry.BoundType == ExactBound {
		return true, entry.Evaluation, entry.BestMove
	} else if entry.BoundType == LowerBound && entry.Evaluation >= beta {
		return true, entry.Evaluation, entry.BestMove
	} else if entry.BoundType == UpperBound && entry.Evaluation <= alpha {
		return true, entry.Evaluation, entry.BestMove
	}

	return false, NoWin, game.Move{}
}

func (t *transpositionTable) StoreHash(hash uint64, evaluation, depth int, entryType BoundType, bestMove game.Move) {
	existing, exists := t.entries[hash]

	// Only replace if the new depth is greater, or it's a new entry
	if !exists || depth > existing.Depth {
		t.entries[hash] = TTEntry{Evaluation: evaluation, Depth: depth, BoundType: entryType, BestMove: bestMove}
	}
}
