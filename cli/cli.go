package cli

import (
	"bufio"
	"fmt"
	"gibhub.com/bef1993/gobblet-gobblers/ai"
	"gibhub.com/bef1993/gobblet-gobblers/game"
	"os"
	"strconv"
	"strings"
)

func PlayGame(human game.Player) {
	board := game.NewBoard()
	var winner game.Player

	printBoard(board)
	for {

		if board.ActivePlayer == human {
			makeHumanMove(board)
		} else {
			makeAIMove(board)
		}
		printBoard(board)

		winner = board.CheckWin()
		if winner != game.None {
			break
		}
	}

	fmt.Printf("Winner: Player %v\n", winner)
}

func makeHumanMove(board *game.Board) {
	var move game.Move
	for {
		move = getHumanMove(board.ActivePlayer)
		fmt.Printf("Playing move %+v\n", move)
		err := board.MakeMove(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
}

func getHumanMove(activePlayer game.Player) game.Move {
	var toRow, toCol int
	var fromRow, fromCol *int
	var size game.Size
	// TODO improve move input logic
	fmt.Println("Enter target coordinates as 'row col'")
	for {
		_, err := fmt.Scan(&toRow, &toCol)
		if err != nil {
			fmt.Println("Invalid input. Please enter two integers.")
			continue
		}
		break
	}

	fmt.Println("Enter piece size (Small = 0, Medium = 1, Large = 2)")
	for {
		_, err := fmt.Scan(&size)
		if err != nil {
			fmt.Println("Invalid piece size. Please enter a valid size.")
			continue
		}
		break
	}

	fmt.Println("Hit Enter to place a new piece or enter coordinates as 'row col' to move existing piece")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			break
		}

		// Split input by whitespace
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid format! Enter as row,col (e.g.  1 2).")
			continue
		}

		// Convert input to integers
		var err1, err2 error
		fromRowInput, err1 := strconv.Atoi(parts[0])
		fromColInput, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input!")
			continue
		}
		fromRow = &fromRowInput
		fromCol = &fromColInput

		break
	}

	move := game.Move{Piece: game.Piece{Size: size, Owner: activePlayer}, To: game.Position{Row: toRow, Col: toCol}}
	if fromRow != nil && fromCol != nil {
		move.From = &game.Position{Row: *fromRow, Col: *fromCol}
	}

	return move
}

func makeAIMove(board *game.Board) {
	move := ai.GetBestMove(board)
	fmt.Printf("Playing move %+v\n", move)
	board.MustMakeMove(move)
}

func DetermineHumanPlayer() (game.Player, error) {
	for {
		var player int
		_, err := fmt.Scanln(&player)
		if err != nil {
			return game.None, err
		}
		if player == 1 {
			return game.Player1, nil
		} else if player == 2 {
			return game.Player2, nil
		} else {
			fmt.Println("Type '1' or '2'")
			continue
		}
	}
}

func printBoard(board *game.Board) {
	fmt.Println("  0   1   2 ")
	fmt.Println(" ───────────")
	// TODO fix misalignment of columns
	for row := 0; row < 3; row++ {
		fmt.Print(row, "| ")
		for col := 0; col < 3; col++ {
			topPiece := board.TopPiece(game.Position{Row: row, Col: col})
			if topPiece == nil {
				fmt.Print("  ")
			} else {
				fmt.Print(topPiece.String() + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println(" ───────────")
}
