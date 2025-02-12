package game

type Move struct {
	Piece Piece
	From  Position
	To    Position
}

type Position struct {
	Row int
	Col int
}

func NewMove(player Player, row, col int, size Size) Move {
	if player == None {
		panic("player must not be None")
	}
	return Move{Piece: Piece{Owner: player, Size: size},
		From: Position{},
		To:   Position{Row: row, Col: col}}
}

func NewMoveExisting(fromRow, fromCol, toRow, toCol int) Move {
	return Move{Piece: Piece{},
		From: Position{Row: fromRow, Col: fromCol},
		To:   Position{Row: toRow, Col: toCol}}
}

func (m Move) IsOutOfBounds() bool {
	return m.From.IsOutOfBounds() || m.To.IsOutOfBounds()
}

func (m Move) MovesExistingPiece() bool {
	return m.Piece == Piece{}
}

func (m Move) PlacesNewPiece() bool {
	return !m.MovesExistingPiece()
}

func (p Position) IsWithinBounds() bool {
	if p.Col < 0 || p.Col > 2 || p.Row < 0 || p.Row > 2 {
		return false
	}
	return true
}

func (p Position) IsOutOfBounds() bool {
	return !p.IsWithinBounds()
}
