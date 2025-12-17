package cli

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"gibhub.com/bef1993/gobblet-gobblers/ai"
	"gibhub.com/bef1993/gobblet-gobblers/game"
)

const RegexMovePattern = "^[abcABC][1-3] ([sS]|[mM]|[lL]|[abcABC][1-3])$"

func PlayGame(human game.Player, maxDepth int) {
	board := game.NewBoard()
	minimax := ai.NewMinimax()
	var winner game.Player

	printBoard(board)
	fmt.Println("Placing piece: b2 S   -   Moving piece: b1 a0")

	for {

		if board.ActivePlayer == human {
			printAvailablePieces(board)
			makeHumanMove(board)
		} else {
			fmt.Println("Waiting for AI to make move ...")
			makeAIMove(board, minimax, maxDepth)
		}
		printBoard(board)

		winner = board.CheckWin()
		if winner != game.None {
			break
		}
	}

	fmt.Printf("Winner: Player %v\n", winner)
	// Wait for a single key press before exiting

	// Consume any leftover input (to prevent immediate exit)
	_, _ = fmt.Scanln()
	_, _ = fmt.Scanln()
}

func makeHumanMove(board *game.Board) {
	var move game.Move
	for {
		move = getHumanMove(board)
		err := board.MakeMove(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
}

func getHumanMove(board *game.Board) (move game.Move) {

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

func makeAIMove(board *game.Board, minimax ai.Minimax, maxDepth int) {
	move := minimax.GetBestMove(board, maxDepth)
	fmt.Printf("AI Move: %v\n", MoveString(move))
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
	for row := 0; row < 3; row++ {
		fmt.Print(row+1, "| ")
		for col := 0; col < 3; col++ {
			topPiece := board.Get(row, col).TopPiece()
			if topPiece == nil {
				fmt.Print(".")
			} else {
				fmt.Print(topPiece.String())
			}
			fmt.Print(" ")
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
	var to *game.Position
	var piece game.Piece

	if moveIsPlacingNewPiece(inputs) {
		to = board.Get(parseCoords(inputs[0]))
		size := letterToSize(inputs[1][0])
		piece = game.Piece{Owner: board.ActivePlayer, Size: size}
	} else {
		from = board.Get(parseCoords(inputs[0]))
		to = board.Get(parseCoords(inputs[1]))
	}

	return game.Move{Piece: piece, From: from, To: to}, nil
}

func parseCoords(input string) (row, col int) {
	row = int(input[1]) - '0' - 1
	col = letterToColIndex(input[0])
	return row, col
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

func sizeToLetter(size game.Size) string {
	switch size {
	case game.Small:
		return "S"
	case game.Medium:
		return "M"
	case game.Large:
		return "L"
	default:
		return "?"
	}
}

func colIndexToLetter(col int) string {
	switch col {
	case 0:
		return "a"
	case 1:
		return "b"
	case 2:
		return "c"
	default:
		return "?"
	}
}

func moveIsPlacingNewPiece(inputs []string) bool {
	return len(inputs[1]) == 1
}

func MoveString(move game.Move) string {
	if move.PlacesNewPiece() {
		return fmt.Sprintf("%v %v", PositionString(move.To), sizeToLetter(move.Piece.Size))
	} else {
		return fmt.Sprintf("%v %v", PositionString(move.From), PositionString(move.To))
	}
}

func PositionString(p *game.Position) string {
	return fmt.Sprintf("%v%v", colIndexToLetter(p.Col), p.Row+1)
}
