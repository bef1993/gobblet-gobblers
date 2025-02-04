package game

type Board struct {
	Grid [3][3][]Piece
}

func (b *Board) PlacePiece(row, col int, piece Piece) {
    b.Grid[row][col] = append(b.Grid[row][col], piece)
}

func (b *Board) TopPiece(row, col int) (*Piece) {
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
