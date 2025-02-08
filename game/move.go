package game

type Move struct {
	Piece Piece
	From  *Position
	To    Position
}

type Position struct {
	Row int
	Col int
}

func NewMove(player Player, row, col int, size Size) Move {
	return Move{Piece: Piece{Owner: player, Size: size},
		From: nil,
		To:   Position{Row: row, Col: col}}
}

func NewMoveExisting(player Player, fromRow, fromCol int, size Size, toRow, toCol int) Move {
	return Move{Piece: Piece{Owner: player, Size: size},
		From: &Position{Row: fromRow, Col: fromCol},
		To:   Position{Row: toRow, Col: toCol}}
}
