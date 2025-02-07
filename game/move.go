package game

type Move struct {
	Piece Piece
	From *Position
	To Position
}

type Position struct {
	Row int
	Col int
}