package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PlayGame(human Player) {
	board := NewBoard()
	activePlayer := Player1
	var winner Player

	for {
		if activePlayer == human {
			makeHumanMove(board, human)
		} else {
			makeAIMove(board, activePlayer)
		}

		if activePlayer == Player1 {
			activePlayer = Player2
		} else {
			activePlayer = Player1
		}

		winner = board.CheckWin()
		if winner != None {
			break
		}
	}

	fmt.Printf("Winner: Player %v\n", winner)
}

func makeHumanMove(board *Board, activePlayer Player) {
	var move Move
	for {
		move = getHumanMove(activePlayer)
		fmt.Printf("Playing move %+v\n", move)
		err := board.MakeMove(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

}

func getHumanMove(activePlayer Player) Move {
	var toRow, toCol int
	var fromRow, fromCol *int
	var size Size
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

	move := Move{Piece: Piece{Size: size, Owner: activePlayer}, To: Position{Row: toRow, Col: toCol}}
	if fromRow != nil && fromCol != nil {
		move.From = &Position{Row: *fromRow, Col: *fromCol}
	}

	return move
}

func makeAIMove(board *Board, activePlayer Player) {
	// TODO make actual ai move
	move := board.GetPossibleMoves(activePlayer)[0]
	fmt.Printf("AI move %+v\n", move)
	board.MustMakeMove(move)
}

func DetermineHumanPlayer() (Player, error) {
	for {
		var player int
		_, err := fmt.Scanln(&player)
		if err != nil {
			return None, err
		}
		if player == 1 {
			return Player1, nil
		} else if player == 2 {
			return Player2, nil
		} else {
			fmt.Println("Type '1' or '2'")
			continue
		}
	}
}
