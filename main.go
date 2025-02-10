package main

import (
	"fmt"
	"gibhub.com/bef1993/gobblet-gobblers/cli"
	"log"
)

func main() {
	fmt.Println("Welcome to Gobblet Gobblers")
	fmt.Println("Do you want to play as Player 1 or Player 2?")
	player, err := cli.DetermineHumanPlayer()
	if err != nil {
		log.Fatal(err)
		return
	}

	cli.PlayGame(player)
}
