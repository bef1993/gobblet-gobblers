package game

import "errors"

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

func (b *Board) PlacePiece(row, col int, piece Piece) error {
	// Check if move is valid
	if (!b.IsValidMove(Move{Row: row, Col: col, Piece: piece})) {
		return errors.New("attempting to make illegal move")
	}
	b.Grid[row][col] = append(b.Grid[row][col], piece)
	b.RemainingPieces[piece.Owner][piece.Size]--

	return nil
}

func (b *Board) TopPiece(row, col int) *Piece {
	stack := b.Grid[row][col]
	if len(stack) == 0 {
		return nil // No pieces in this cell
	}
	return &stack[len(stack)-1] // Return top piece
}

func (b *Board) CheckWin() Player {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if winner := b.checkLine(b.TopPiece(i, 0), b.TopPiece(i, 1), b.TopPiece(i, 2)); winner != None {
			return winner
		}
		if winner := b.checkLine(b.TopPiece(0, i), b.TopPiece(1, i), b.TopPiece(2, i)); winner != None {
			return winner
		}
	}

	// Check diagonals
	if winner := b.checkLine(b.TopPiece(0, 0), b.TopPiece(1, 1), b.TopPiece(2, 2)); winner != None {
		return winner
	}
	if winner := b.checkLine(b.TopPiece(0, 2), b.TopPiece(1, 1), b.TopPiece(2, 0)); winner != None {
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

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			for _, piece := range b.AvailablePieces(player) {
				move := Move{Row: row, Col: col, Piece: piece}
				if b.IsValidMove(move) {
					moves = append(moves, move)
				}
			}
		}
	}

	// TODO implement moving of already placed pieces

	return moves
}

func (b *Board) IsValidMove(move Move) bool {
	row, col, piece := move.Row, move.Col, move.Piece

	// Check bounds
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return false
	}

	// Check if player has piece available
	if b.RemainingPieces[piece.Owner][piece.Size] < 1 {
		return false
	}

	stack := b.Grid[row][col]

	// If empty, any piece can be placed
	if len(stack) == 0 {
		return true
	}

	// Get the top piece in the stack
	topPiece := stack[len(stack)-1]

	// Ensure the piece is larger then the already placed piece
	return piece.Size > topPiece.Size
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
