package game

import (
	"math/rand"
)

var zobristTable [3][3][6]uint64 // 3x3 board, 6 unique pieces
var activePlayerHash [3]uint64   // 1 random value per player

// InitZobrist initializes the hash table
func InitZobrist() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			for piece := 0; piece < 6; piece++ {
				zobristTable[row][col][piece] = rand.Uint64()
			}
		}
	}
	activePlayerHash[1] = rand.Uint64() // Player 1 hash
	activePlayerHash[2] = rand.Uint64() // Player 2 hash
}

func init() {
	InitZobrist() // Runs automatically before `main()`
}

func GetZobristValue(position *Position, piece Piece) uint64 {
	return zobristTable[position.Row][position.Col][piece.ID()]
}

func GetPlayerZobristValue(activePlayer Player) uint64 {
	return activePlayerHash[activePlayer]
}
