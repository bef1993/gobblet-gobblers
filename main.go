package main

import (
	"flag"
	"fmt"
	"log"

	"gibhub.com/bef1993/gobblet-gobblers/cli"
)

func main() {
	maxDepth := flag.Int("maxDepth", 8, "the maximum search depth for the AI")
	flag.Parse()

	fmt.Println("Welcome to Gobblet Gobblers")
	fmt.Println("Do you want to play as Player 1 or Player 2?")
	player, err := cli.DetermineHumanPlayer()
	if err != nil {
		log.Fatal(err)
		return
	}

	cli.PlayGame(player, *maxDepth)
}
