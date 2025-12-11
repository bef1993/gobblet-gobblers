package game

import (
	"errors"
	"fmt"
)

type Board struct {
	Grid            [3][3]Position
	Lines           [8]Line
	RemainingPieces map[Player][]int
	ActivePlayer    Player
	Hash            uint64
}

type Line [3]*Position

type Position struct {
	Row    int
	Col    int
	Pieces []Piece
}

func NewBoard() *Board {
	board := &Board{
		RemainingPieces: map[Player][]int{
			Player1: {2, 2, 2},
			Player2: {2, 2, 2},
		},
		ActivePlayer: Player1,
		Hash:         GetPlayerZobristValue(Player1),
	}
	board.initializePositions()
	board.initializeLines()
	return board
}

func (b *Board) Get(row, col int) *Position {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		panic("can not get position out of bounds")
	}
	return &b.Grid[row][col]
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

	if move.MovesExistingPiece() {
		pieceToMove := move.From.TopPiece()
		b.removePiece(move.From)
		b.placePiece(move.To, *pieceToMove)
	} else {
		b.placePiece(move.To, move.Piece)

	}

	b.switchActivePlayer()
	return nil
}

func (b *Board) MustUndoMove(move Move) {
	if move.MovesExistingPiece() {
		pieceToMove := move.To.TopPiece()
		b.removePiece(move.To)
		b.placePiece(move.From, *pieceToMove)
	} else {
		b.removePiece(move.To)
	}

	b.switchActivePlayer()
}

func (b *Board) placePiece(p *Position, piece Piece) {
	b.Hash ^= GetZobristValue(p, piece)
	p.Pieces = append(p.Pieces, piece)
	b.RemainingPieces[piece.Owner][piece.Size]--
}

func (b *Board) removePiece(p *Position) {
	topPiece := p.TopPiece()
	if topPiece == nil {
		panic("attempting to remove piece from empty position")
	}
	b.Hash ^= GetZobristValue(p, *topPiece)
	p.Pieces = p.Pieces[:len(p.Pieces)-1]
	b.RemainingPieces[topPiece.Owner][topPiece.Size]++
}

func (p Position) TopPiece() *Piece {
	if len(p.Pieces) == 0 {
		return nil // No pieces in this cell
	}
	// Create and return a copy of the top piece
	topPiece := p.Pieces[len(p.Pieces)-1]
	return &topPiece
}

func (b *Board) CheckWin() Player {
	for _, line := range b.Lines {
		if winner := line.CheckWin(); winner != None {
			return winner
		}
	}
	return None // No winner yet
}

func (l Line) CheckWin() Player {
	p1 := l[0].TopPiece()
	p2 := l[1].TopPiece()
	p3 := l[2].TopPiece()

	if p1 == nil || p2 == nil || p3 == nil {
		return None // At least one empty cell, no win
	}
	if p1.Owner == p2.Owner && p2.Owner == p3.Owner {
		return p1.Owner // Return winning player
	}
	return None
}

func (b *Board) GetPossibleMoves() []Move {
	var moves []Move

	// Try placing new pieces
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			for _, size := range b.AvailablePieceSizes(b.ActivePlayer) {
				move := NewMove(b.ActivePlayer, b.Get(row, col), size)
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
			fromPos := b.Get(fromRow, fromCol)
			topPiece := fromPos.TopPiece()

			// Skip empty positions
			if topPiece == nil {
				continue
			}

			// Only consider moves where the top piece belongs to the player
			if topPiece.Owner != b.ActivePlayer {
				continue
			}

			// Try moving to every other position
			for toRow := 0; toRow < 3; toRow++ {
				for toCol := 0; toCol < 3; toCol++ {
					toPos := b.Get(toRow, toCol)

					// Don't move to the same position
					if fromPos == toPos {
						continue
					}

					move := NewMoveExisting(fromPos, toPos)
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

	// Check if positions exist within grid
	if to == nil || b.Get(to.Row, to.Col) != to {
		return false, errors.New("TO position not found in grid")
	}
	if from != nil && &b.Grid[from.Row][from.Col] != from {
		return false, errors.New("FROM position not found in grid")
	}

	// Check winner
	if b.CheckWin() != None {
		return false, errors.New(fmt.Sprintf("game is already won"))
	}

	// Check piece belongs to active player
	if move.PlacesNewPiece() && move.Piece.Owner != b.ActivePlayer {
		return false, errors.New("piece does not belong to active player")
	}

	// Check if player has piece available
	if move.PlacesNewPiece() && !b.hasPieceAvailable(piece) {
		return false, errors.New("player does not have piece available")
	}

	// If Move moves a placed piece
	if move.MovesExistingPiece() {

		// Check that a piece exists on position
		pieceToMove := from.TopPiece()
		if pieceToMove == nil {
			return false, errors.New(fmt.Sprintf("piece does not exist on position"))
		}

		// Check that piece from position belongs to active player
		if pieceToMove.Owner != b.ActivePlayer {
			return false, errors.New(fmt.Sprintf("piece does not belong to active player"))
		}

		//From and To positions must be different
		if to == from {
			return false, errors.New("origin and target position are identical")
		}

		// Check that moving the piece would not cause the other player to win
		originalStack := from.Pieces
		// Temporarily remove the top piece
		from.Pieces = originalStack[:len(originalStack)-1]
		// Check if the opponent wins
		winner := b.CheckWin()
		// Restore board state
		from.Pieces = originalStack
		// If opponent wins, this move is invalid
		if winner != None {
			return false, errors.New("moving piece would cause the other player to win")
		}
	}

	// If empty, any piece can be placed
	if len(to.Pieces) == 0 {
		return true, nil
	}

	// Use pieceToMove if necessary
	if piece == (Piece{}) {
		piece = *from.TopPiece()
	}

	// Ensure the piece is larger than the already placed piece
	if piece.Size <= to.TopPiece().Size {
		return false, errors.New("piece not larger than existing piece on position")
	}

	return true, nil
}

func (b *Board) AvailablePieceSizes(player Player) (sizes []Size) {
	if b.RemainingPieces[player][Small] > 0 {
		sizes = append(sizes, Small)
	}
	if b.RemainingPieces[player][Medium] > 0 {
		sizes = append(sizes, Medium)
	}
	if b.RemainingPieces[player][Large] > 0 {
		sizes = append(sizes, Large)
	}

	return sizes
}

func (b *Board) hasPieceAvailable(piece Piece) bool {
	return b.RemainingPieces[piece.Owner][piece.Size] >= 1
}

func (b *Board) switchActivePlayer() {
	b.Hash ^= GetPlayerZobristValue(b.ActivePlayer)
	b.ActivePlayer = b.ActivePlayer.Opponent()
	b.Hash ^= GetPlayerZobristValue(b.ActivePlayer)
}

func (b *Board) initializePositions() {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			b.Grid[r][c] = Position{Row: r, Col: c, Pieces: []Piece{}}
		}
	}
}

func (b *Board) initializeLines() {
	// Rows
	for r := 0; r < 3; r++ {
		b.Lines[r] = Line{&b.Grid[r][0], &b.Grid[r][1], &b.Grid[r][2]}
	}

	// Columns
	for c := 0; c < 3; c++ {
		b.Lines[c+3] = Line{
			&b.Grid[0][c], &b.Grid[1][c], &b.Grid[2][c],
		}
	}

	// Diagonals
	b.Lines[6] = Line{&b.Grid[0][0], &b.Grid[1][1], &b.Grid[2][2]}
	b.Lines[7] = Line{&b.Grid[0][2], &b.Grid[1][1], &b.Grid[2][0]}
}
