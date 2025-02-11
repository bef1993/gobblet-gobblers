package cli

import (
	"errors"
	"fmt"
	"gibhub.com/bef1993/gobblet-gobblers/ai"
	"gibhub.com/bef1993/gobblet-gobblers/game"
	"regexp"
	"strings"
	"unicode"
)

const RegexMovePattern = "^[abcABC][1-3] ([sS]|[mM]|[lL]|[abcABC][1-3])$"

func PlayGame(human game.Player) {
	board := game.NewBoard()
	var winner game.Player

	printBoard(board)
	for {

		if board.ActivePlayer == human {
			printAvailablePieces(board)
			makeHumanMove(board)
		} else {
			fmt.Println("Waiting for AI to make move ...")
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
		move = getHumanMove(board)
		fmt.Printf("Playing move %+v\n", move)
		err := board.MakeMove(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
}

func getHumanMove(board *game.Board) (move game.Move) {
	fmt.Println("Placing piece: b2 S")
	fmt.Println("Moving piece: b1 a0")
	fmt.Println("Enter your move:")
	for {
		var input1, input2 string
		n, err := fmt.Scan(&input1, &input2)
		if err != nil || n != 2 {
			fmt.Println("Invalid input")
			continue
		}
		move, err = ParseMove(strings.Join([]string{input1, input2}, " "), board)
		if err != nil {
			fmt.Println("Invalid input. Please enter move again:")
			continue
		}
		return move
	}

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
	fmt.Println("  a   b   c ")
	fmt.Println(" ───────────")
	// TODO fix misalignment of columns
	for row := 0; row < 3; row++ {
		fmt.Print(row+1, "| ")
		for col := 0; col < 3; col++ {
			topPiece := board.TopPiece(game.Position{Row: row, Col: col})
			if topPiece == nil {
				fmt.Print(". ")
			} else {
				fmt.Print(topPiece.String() + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println(" ───────────")
}

func printAvailablePieces(board *game.Board) {
	activePlayer := board.ActivePlayer
	fmt.Printf("Available pieces: %v Small, %v Medium, %v Large\n",
		board.RemainingPieces[activePlayer][game.Small],
		board.RemainingPieces[activePlayer][game.Medium],
		board.RemainingPieces[activePlayer][game.Large])

}

func ParseMove(input string, board *game.Board) (game.Move, error) {
	matched, err := regexp.Match(RegexMovePattern, []byte(input))
	if err != nil {
		panic("invalid regex pattern")
	}
	if !matched {
		return game.Move{}, errors.New("invalid move input")
	}

	inputs := strings.Split(input, " ")
	var from *game.Position
	var to game.Position
	var piece game.Piece
	if moveIsPlacingNewPiece(inputs) {
		to = parsePosition(inputs[0])
		size := letterToSize(inputs[1][0])
		piece = game.Piece{Owner: board.ActivePlayer, Size: size}
	} else {
		position := parsePosition(inputs[0])
		from = &position
		to = parsePosition(inputs[1])
		topPiece := board.TopPiece(*from)
		if topPiece == nil {
			piece = game.Piece{}
		} else {
			piece = *topPiece
		}
	}

	return game.Move{Piece: piece, From: from, To: to}, nil
}

func parsePosition(input string) game.Position {
	row := int(input[1]) - '0' - 1
	col := letterToColIndex(input[0])
	return game.Position{Row: row, Col: col}
}

func letterToColIndex(letter uint8) int {
	switch unicode.ToLower(rune(letter)) {
	case 'a':
		return 0
	case 'b':
		return 1
	case 'c':
		return 2
	default:
		panic("invalid col letter")
	}
}

func letterToSize(letter uint8) game.Size {
	switch unicode.ToLower(rune(letter)) {
	case 's':
		return game.Small
	case 'm':
		return game.Medium
	case 'l':
		return game.Large
	default:
		panic("invalid size letter")
	}
}

func moveIsPlacingNewPiece(inputs []string) bool {
	return len(inputs[1]) == 1
}
