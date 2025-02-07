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
