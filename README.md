# TicTacGo: Tic-Tac-Toe CLI in Go

A simple command-line Tic-Tac-Toe game written in Go. This project is designed for Go beginners to understand Go basics, project structure, and CLI interaction.

---

## Features
- Play Tic-Tac-Toe in your terminal
- Two-player mode (play with a friend)
- Simple, readable Go code
- Modular structure for easy learning

---

## Project Structure
```
.
├── main.go           # Entry point for the CLI game
├── go.mod            # Go module definition
├── game/
│   ├── board.go      # Board logic (printing, marking, win checks)
│   ├── game.go       # Game loop and player turns
│   └── player.go     # Player struct
└── .vscode/
    └── launch.json   # VSCode debug configuration
```

---

## Getting Started

### Prerequisites
- [Go](https://go.dev/dl/) 1.24 or newer

### Clone the Repository
```sh
git clone <your-repo-url>
cd tictacgo
```

### Run the Game
You can run the game directly:
```sh
go run main.go
```

Or build an executable:
```sh
go build
./tictacgo   # or tictacgo.exe on Windows
```

### VSCode Debugging
A sample `.vscode/launch.json` is included. You can debug the game by selecting "Launch Package" in the VSCode Run & Debug panel.

---

## How to Play
- The game is played on a 3x3 grid.
- Two players take turns: Player 1 (X) and Player 2 (O).
- On your turn, enter the row and column numbers (both 0-based) to place your symbol.
- The first player to get three of their symbols in a row (vertically, horizontally, or diagonally) wins.
- If all cells are filled and no player has won, the game ends in a draw.

### Example Game
```
Rounik (X), enter row and column: 0 0

X  -  -  
-  -  -  
-  -  -  

Rohan (O), enter row and column: 1 1

X  -  -  
-  O  -  
-  -  -  

... (game continues)
```

---

## Customization
- Change player names or symbols in `main.go`.
- Modify board size or add features by editing the files in the `game/` directory.

---

## License
This project is open source and available under the [MIT License](LICENSE) (add a LICENSE file if needed).
