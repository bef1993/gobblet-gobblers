package game

type Move struct {
	Piece Piece
	From  *Position
	To    *Position
}

func NewMove(player Player, to *Position, size Size) Move {
	if player == None {
		panic("player must not be None")
	}
	if to == nil {
		panic("to must not be nil")
	}

	return Move{
		Piece: Piece{Owner: player, Size: size},
		From:  nil,
		To:    to,
	}
}

func NewMoveExisting(from, to *Position) Move {
	if from == nil || to == nil {
		panic("from and to must not be nil")
	}

	return Move{
		Piece: Piece{},
		From:  from,
		To:    to,
	}
}

func (m Move) MovesExistingPiece() bool {
	return m.Piece == Piece{}
}

func (m Move) PlacesNewPiece() bool {
	return !m.MovesExistingPiece()
}
