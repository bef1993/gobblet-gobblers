package game

type Size int
type Player int

const (
	None    Player = 0
	Player1 Player = 1
	Player2 Player = 2
)

const (
	Small  Size = 0
	Medium Size = 1
	Large  Size = 2
)

type Piece struct {
	Owner Player
	Size  Size
}

func (p Player) Opponent() Player {
	switch p {
	case Player1:
		return Player2
	case Player2:
		return Player1
	default:
		return None
	}
}
