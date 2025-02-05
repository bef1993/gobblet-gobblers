package game

type Size int
type Player int

const (
	None    = 0
	Player1 = 1
	Player2 = 2
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
