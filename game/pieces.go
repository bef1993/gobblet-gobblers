package game

import (
	"github.com/fatih/color"
	"strconv"
)

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

func (piece Piece) String() string {
	var size string
	switch piece.Size {
	case Small:
		size = "S"
	case Medium:
		size = "M"
	case Large:
		size = "L"
	}
	str := strconv.Itoa(int(piece.Owner)) + size
	if piece.Owner == Player1 {
		return color.RedString(str)
	} else {
		return color.GreenString(str)
	}
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
