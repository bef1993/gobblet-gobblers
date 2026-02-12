# Gobblet Gobblers - AI Agent Information

This document provides an overview of the Gobblet Gobblers codebase, specifically tailored for AI agents to understand the project structure, game logic, and AI implementation.

## Project Overview

**Gobblet Gobblers** is a strategic board game implemented in Go. It's a variation of Tic-Tac-Toe where players can gobble opponent's pieces.

### Key Directories

- `game/`: Contains the core game logic (Board, Pieces, Moves, Rules).
- `ai/`: Contains the AI implementation (Minimax, Evaluator, Transposition Table).
- `cli/`: Contains the Command Line Interface for playing the game.
- `main.go`: Entry point of the application.

## Game Logic (`game/`)

### Board Representation (`game/board.go`)

The board is a 3x3 grid represented by the `Board` struct.
- `Grid [3][3]Position`: The 3x3 grid of positions.
- `Position`: Represents a cell on the board. It contains a stack of pieces (`Pieces []Piece`).
- `RemainingPieces`: Tracks the pieces available for each player to place.
- `ActivePlayer`: The player whose turn it is.
- `Hash`: Zobrist hash of the current board state.

### Pieces (`game/pieces.go`)

- `Piece`: Has an `Owner` (Player1 or Player2) and a `Size` (Small, Medium, Large).
- `Player`: Enum for Player1 (1) and Player2 (2).
- `Size`: Enum for Small (0), Medium (1), Large (2).

### Moves (`game/move.go`)

Moves are represented by the `Move` struct.
- `NewMove(player, to, size)`: Creates a move to place a new piece.
- `NewMoveExisting(from, to)`: Creates a move to move an existing piece on the board.
- `IsValidMove(move)`: Validates a move according to game rules.

### Rules

1.  **Placement**: Players can place a new piece on an empty square or gobble a smaller piece.
2.  **Movement**: Players can move their own exposed pieces to another square (empty or occupied by a smaller piece).
3.  **Winning**: Align 3 pieces of your color in a row, column, or diagonal.
4.  **Special Rule**: You cannot move a piece if it uncovers an opponent's piece that completes a line for them (this is checked in `IsValidMove`).

## AI Implementation (`ai/`)

The AI uses the Minimax algorithm with Alpha-Beta pruning.

### Minimax (`ai/minimax.go`)

- `GetBestMove(board, maxDepth)`: Returns the best move for the current board state.
- `minimax(...)`: Recursive function implementing the algorithm.
- **Optimizations**:
    - **Alpha-Beta Pruning**: Prunes branches that don't need to be explored.
    - **Transposition Table**: Caches board evaluations to avoid re-calculating identical states.
    - **Move Ordering**: Sorts moves based on a heuristic to improve pruning efficiency.

### Evaluator (`ai/evaluator.go`)

- `Evaluate(board, depth)`: Returns a score for the board state. Positive for Player 1 advantage, negative for Player 2.
- `EvaluateMove(board, move)`: Heuristic to score a single move for sorting.

### Transposition Table (`ai/transposition_table.go`)

- Stores evaluations of board states using Zobrist hashing (`game/zobrist_hash.go`).

## Running the Game

- Build: `make build`
- Run: `make run` or `./gobblet_gobblers`
- Test: `make test`

## CLI Usage

When running the game, you can choose to play as Player 1 or Player 2.
Moves are entered in the format:
- Place: `[position] [size]` (e.g., `b2 S`)
- Move: `[from] [to]` (e.g., `a1 c3`)
