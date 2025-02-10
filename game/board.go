package game

import (
	"errors"
	"fmt"
)

type Board struct {
	Grid            [3][3][]Piece
	RemainingPieces map[Player][]int
}

func NewBoard() *Board {
	return &Board{
		RemainingPieces: map[Player][]int{
			Player1: {2, 2, 2},
			Player2: {2, 2, 2},
		},
	}
}

func (b *Board) MustMakeMove(move Move) {
	err := b.MakeMove(move)
	if err != nil {
		panic(err)
	}
}

func (b *Board) MakeMove(move Move) error {
	valid, err := b.IsValidMove(move)
	if !valid {
		return err
	}

	if move.From != nil {
		b.removePiece(*move.From)
	}
	b.placePiece(move.To, move.Piece)

	return nil
}

func (b *Board) MustUndoMove(move Move) {
	// TODO add checks if valid
	if move.From != nil {
		b.placePiece(*move.From, move.Piece)
	}
	b.removePiece(move.To)
}

func (b *Board) placePiece(p Position, piece Piece) {
	stack := b.getPositionStack(p)
	b.Grid[p.Row][p.Col] = append(stack, piece)
	b.RemainingPieces[piece.Owner][piece.Size]--
}

func (b *Board) removePiece(p Position) {
	topPiece := b.TopPiece(p)
	if topPiece == nil {
		panic("attempting to remove piece from empty position")
	}
	stack := b.getPositionStack(p)
	b.Grid[p.Row][p.Col] = stack[:len(stack)-1]
	b.RemainingPieces[topPiece.Owner][topPiece.Size]++
}

func (b *Board) TopPiece(p Position) *Piece {
	stack := b.Grid[p.Row][p.Col]
	if len(stack) == 0 {
		return nil // No pieces in this cell
	}
	return &stack[len(stack)-1] // Return top piece
}

func (b *Board) CheckWin() Player {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if winner := b.checkLine(b.TopPiece(Position{i, 0}), b.TopPiece(Position{i, 1}), b.TopPiece(Position{i, 2})); winner != None {
			return winner
		}
		if winner := b.checkLine(b.TopPiece(Position{0, i}), b.TopPiece(Position{1, i}), b.TopPiece(Position{2, i})); winner != None {
			return winner
		}
	}

	// Check diagonals
	if winner := b.checkLine(b.TopPiece(Position{0, 0}), b.TopPiece(Position{1, 1}), b.TopPiece(Position{2, 2})); winner != None {
		return winner
	}
	if winner := b.checkLine(b.TopPiece(Position{0, 2}), b.TopPiece(Position{1, 1}), b.TopPiece(Position{2, 0})); winner != None {
		return winner
	}

	return None // No winner yet
}

func (b *Board) checkLine(p1, p2, p3 *Piece) Player {
	if p1 == nil || p2 == nil || p3 == nil {
		return None // At least one empty cell, no win
	}
	if p1.Owner == p2.Owner && p2.Owner == p3.Owner {
		return p1.Owner // Return winning player
	}
	return None
}

func (b *Board) GetPossibleMoves(player Player) []Move {
	var moves []Move

	// Try placing new pieces
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			for _, piece := range b.AvailablePieces(player) {
				move := Move{Piece: piece, From: nil, To: Position{row, col}}
				valid, _ := b.IsValidMove(move)
				if valid {
					moves = append(moves, move)
				}
			}
		}
	}

	//Try moving already placed pieces
	for fromRow := 0; fromRow < 3; fromRow++ {
		for fromCol := 0; fromCol < 3; fromCol++ {
			fromPos := Position{Row: fromRow, Col: fromCol}
			stack := b.getPositionStack(fromPos)
			topPiece := b.TopPiece(fromPos)

			// Skip empty positions
			if len(stack) == 0 {
				continue
			}

			// Only consider moves where the top piece belongs to the player
			if topPiece.Owner != player {
				continue
			}

			// Try moving to every other position
			for toRow := 0; toRow < 3; toRow++ {
				for toCol := 0; toCol < 3; toCol++ {
					toPos := Position{Row: toRow, Col: toCol}

					// Don't move to the same position
					if fromPos == toPos {
						continue
					}

					move := Move{Piece: *topPiece, From: &fromPos, To: toPos}
					if valid, _ := b.IsValidMove(move); valid {
						moves = append(moves, move)
					}
				}
			}
		}
	}
	return moves
}

func (b *Board) IsValidMove(move Move) (bool, error) {
	from, to, piece := move.From, move.To, move.Piece

	// Check bounds
	if to.IsOutOfBounds() {
		return false, errors.New(fmt.Sprintf("move is out of bounds: %+v", to))
	}

	// Check if player has piece available
	if from == nil && !b.hasPieceAvailable(piece) {
		return false, errors.New(fmt.Sprintf("no remaining piece: %+v", piece))
	}

	// If Move moves a placed piece
	if from != nil {
		// Check bounds
		if from.IsOutOfBounds() {
			return false, errors.New(fmt.Sprintf("trying to move from out of bounds: %+v", from))
		}

		// Check that piece is actually there
		pieceOnStack := b.TopPiece(*from)
		if pieceOnStack == nil || piece != *pieceOnStack {
			return false, errors.New(fmt.Sprintf("trying to move a non-existing piece: %+v %+v", from, piece))
		}

		//From and To positions must be different
		if to == *from {
			return false, errors.New(fmt.Sprintf("trying to move a piece to the same location: %+v", from))
		}

		// Check that moving the piece would not cause the other player to win
		originalStack := b.getPositionStack(*from)
		// Temporarily remove the top piece
		b.Grid[from.Row][from.Col] = originalStack[:len(originalStack)-1]
		// Check if the opponent wins
		winner := b.CheckWin()
		// Restore board state
		b.Grid[from.Row][from.Col] = originalStack
		// If opponent wins, this move is invalid
		if winner != None {
			return false, errors.New(fmt.Sprintf("moving this piece would cause the other player to win: %+v", from))
		}
	}

	stack := b.Grid[to.Row][to.Col]

	// If empty, any piece can be placed
	if len(stack) == 0 {
		return true, nil
	}

	// Get the top piece in the stack
	topPiece := stack[len(stack)-1]

	// Ensure the piece is larger than the already placed piece
	if piece.Size <= topPiece.Size {
		return false, errors.New(fmt.Sprintf("piece to place %+v must be larger than top piece %+v", piece, topPiece))
	}

	return true, nil
}

func (b *Board) AvailablePieces(player Player) (pieces []Piece) {
	if b.RemainingPieces[player][Small] > 0 {
		pieces = append(pieces, Piece{Owner: player, Size: Small})
	}
	if b.RemainingPieces[player][Medium] > 0 {
		pieces = append(pieces, Piece{Owner: player, Size: Medium})
	}
	if b.RemainingPieces[player][Large] > 0 {
		pieces = append(pieces, Piece{Owner: player, Size: Large})
	}

	return pieces
}

func (b *Board) hasPieceAvailable(piece Piece) bool {
	return b.RemainingPieces[piece.Owner][piece.Size] >= 1
}

func (b *Board) getPositionStack(p Position) []Piece {
	return b.Grid[p.Row][p.Col]
}
