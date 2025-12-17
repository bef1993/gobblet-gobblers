# Gobblet Gobblers

**Gobblet Gobblers** is a fun and strategic two-player board game that adds a twist to traditional Tic-Tac-Toe. Players take turns placing or moving pieces on a 3×3 grid, but with a unique rule: larger pieces can **"gobble"** smaller ones, covering them up and dynamically changing the board state. The first player to align three of their pieces in a row, column, or diagonal wins.
Each player has two pieces of each size available.

## Move Commands

Moves in Gobblet Gobblers follow this format:

- **Placing a new piece**: Specify the target position and the piece size (S, M, or L).
- **Moving an existing piece**: Specify the starting and destination positions.

### Examples:
Place a Small piece at B2:
```
b2 S
```
Move the piece from A1 to C3
```
a1 c3
```

## Special Rules
A move that would cause the other player to have three of his pieces aligned (by uncovering one of their pieces) is considered illegal.

## AI & Minimax Algorithm

This implementation includes an **AI opponent** powered by the **Minimax algorithm** with **Alpha-Beta pruning**.

### Minimax Optimizations

- **Minimax Algorithm**: The AI evaluates all possible moves to find the optimal play by simulating future game states.
- **Alpha-Beta Pruning**: Optimizes Minimax by eliminating unnecessary branches, making the AI more efficient.
- **Heuristic Move Sorting**: Instead of exploring moves in a random order, the AI first evaluates each possible move with a heuristic function and sorts them. By searching the most promising moves first, the algorithm is much more likely to trigger alpha-beta pruning, leading to a significant performance increase.
- **Incremental Zobrist Hashing**: Used for board state hashing to speed up repeated state evaluations.