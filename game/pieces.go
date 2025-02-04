package game

type Size int
type Player int

const (
    None = 0
    Player1 = 1
    Player2 = 2
)

const (
    Small  Size = 1
    Medium Size = 2
    Large  Size = 3
)

type Piece struct {
    Owner Player
    Size  Size
}
