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

// TODO implement 2 Move constructors (for placing and moving a piece)
