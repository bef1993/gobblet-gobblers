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

func (b *Board) MakeMove(move Move) error {
	if !b.IsValidMove(move) {
		return errors.New("attempting to make illegal move")
	}

	if move.From != nil {
		b.removePiece(*move.From)
	}
	b.placePiece(move.To, move.Piece)

	return nil
}

func (b *Board) placePiece(p Position, piece Piece) {
	stack := b.GetPositionStack(p)
	b.Grid[p.Row][p.Col] = append(stack, piece)
	b.RemainingPieces[piece.Owner][piece.Size]--
}

func (b *Board) removePiece(p Position) {
	if b.TopPiece(p) == nil {
		panic("attempting to remove piece from empty position")
	}
	stack := b.GetPositionStack(p)
	b.Grid[p.Row][p.Col] = stack[:len(stack)-1]
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

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			for _, piece := range b.AvailablePieces(player) {
				move := Move{Piece: piece, From: nil, To: Position{row, col}}
				if b.IsValidMove(move) {
					moves = append(moves, move)
				}
			}
		}
	}

	// TODO should also return moves where a piece is moved

	return moves
}

func (b *Board) IsValidMove(move Move) bool {

	from, to, piece := move.From, move.To, move.Piece

	// Check bounds
	if !isWithinBounds(move.To) {
		return false
	}

	// Check if player has piece available
	if from == nil && !b.hasPieceAvailable(piece) {
		return false
	}

	// If Move moves a placed piece
	if from != nil {
		// Check bounds
		if !isWithinBounds(*move.From) {
			return false
		}

		// Check that piece is actually there
		pieceOnStack := b.TopPiece(*from)
		if pieceOnStack == nil || piece != *pieceOnStack {
			return false // TODO does this actually work?
		}

		//From and To positions must be different
		if to == *from {
			return false
		}

		// Check that moving the piece would not cause the other player to win
		originalStack := b.GetPositionStack(*from)
		// Temporarily remove the top piece
		b.Grid[from.Row][from.Col] = originalStack[:len(originalStack)-1]
		// Check if the opponent wins
		winner := b.CheckWin()
		// Restore board state
		b.Grid[from.Row][from.Col] = originalStack
		// If opponent wins, this move is invalid
		if winner != None {
			return false
		}
	}

	stack := b.Grid[to.Row][to.Col]

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

func isWithinBounds(p Position) bool {
	if p.Col < 0 || p.Col > 2 || p.Row < 0 || p.Row > 2 {
		return false
	}
	return true
}

func (b *Board) hasPieceAvailable(piece Piece) bool {
	return b.RemainingPieces[piece.Owner][piece.Size] >= 1
}

func (b *Board) GetPositionStack(p Position) []Piece {
	return b.Grid[p.Row][p.Col]
}
