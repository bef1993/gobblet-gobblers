package game

import (
	"github.com/fatih/color"
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
	if piece.Owner == Player1 {
		return color.RedString(size)
	} else {
		return color.GreenString(size)
	}
}

func (piece Piece) ID() int {
	if piece.Owner == Player1 {
		return int(piece.Size)
	} else {
		return int(piece.Owner) + 3
	}
}

func (player Player) String() string {
	switch player {
	case None:
		return "None"
	case Player1:
		return "Player 1"
	case Player2:
		return "Player 2"
	default:
		return "Unknown"
	}
}

func (player Player) Opponent() Player {
	switch player {
	case Player1:
		return Player2
	case Player2:
		return Player1
	default:
		panic("can not get opponent of NonePlayer")
	}
}
