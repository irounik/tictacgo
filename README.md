# TicTacGo: Tic-Tac-Toe Game in Go

A Tic-Tac-Toe game written in Go with both command-line interface (CLI) and HTTP server modes. This project demonstrates Go basics, project structure, modular design, and both CLI and web server development.

---

## Features
- **CLI Mode**: Play Tic-Tac-Toe in your terminal
- **Server Mode**: HTTP server for web-based gameplay
- Two-player mode (play with a friend)
- Modular architecture with separate CLI and server packages
- Authentication system for server mode
- Clean, readable Go code structure

---

## Project Structure
```
.
├── main.go              # Entry point with mode selection
├── go.mod               # Go module definition
├── cli/
│   └── cli.go          # CLI game setup and play logic
├── game/
│   ├── board.go        # Board logic (printing, marking, win checks)
│   ├── game.go         # Game loop and player turns
│   └── player.go       # Player struct
├── server/
│   ├── server.go       # HTTP server setup and startup
│   ├── config/
│   │   └── contstrants.go  # Configuration constants
│   ├── handler/
│   │   ├── auth_handler.go # Authentication handlers
│   │   └── game_handler.go # Game-related handlers
│   ├── middleware/
│   │   └── auth_middleware.go # Authentication middleware
│   ├── model/          # Data models
│   ├── repository/     # Data access layer
│   ├── service/        # Business logic layer
│   ├── router/         # HTTP routing
│   └── db/            # Database layer
└── .vscode/
    └── launch.json    # VSCode debug configuration
```

---

## Getting Started

### Prerequisites
- [Go](https://go.dev/dl/) 1.24 or newer

### Clone the Repository
```sh
git clone https://github.com/irounik/tictacgo
cd tictacgo
```

### Running the Application

#### CLI Mode (Default)
Run the command-line version:
```sh
go run main.go
# or explicitly
go run main.go -mode=cli
```

#### Server Mode
Start the HTTP server:
```sh
go run main.go -mode=server -port=8080
```

#### Building Executable
Build the application:
```sh
go build
./tictacgo -mode=cli     # CLI mode
./tictacgo -mode=server  # Server mode
```

### Command Line Options
- `-mode`: Specify game mode (`cli` or `server`)
- `-port`: Port for server mode (default: 8080)

### VSCode Debugging
A sample `.vscode/launch.json` is included. You can debug the game by selecting "Launch Package" in the VSCode Run & Debug panel.

---

## How to Play

### CLI Mode
- The game is played on a 3x3 grid
- Two players take turns: Player 1 (X) and Player 2 (O)
- On your turn, enter the row and column numbers (both 0-based) to place your symbol
- The first player to get three of their symbols in a row (vertically, horizontally, or diagonally) wins
- If all cells are filled and no player has won, the game ends in a draw

### Server Mode
- Start the server with `go run main.go -mode=server`
- Access the game through HTTP endpoints
- Authentication required for game access
- Session-based user management

### Example CLI Game
```
Enter name for first player (X): Rounik
Enter name for second player (O): Rohan

-  -  -  
-  -  -  
-  -  -  

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

## Architecture

### CLI Package
- `cli.go`: Handles game setup and main game loop
- Integrates with the `game` package for core game logic

### Server Package
- **Layered Architecture**: Handler → Service → Repository → Database
- **Authentication**: Session-based auth with middleware
- **Modular Design**: Separate packages for different concerns
- **HTTP Routing**: RESTful API endpoints

### Game Package
- Core game logic shared between CLI and server modes
- Board management, win detection, and game state

---

## Customization
- Modify game logic in the `game/` directory
- Add new CLI features in `cli/cli.go`
- Extend server functionality in the `server/` packages
- Change authentication logic in `server/middleware/` and `server/service/`

---

## Development
The project follows Go best practices:
- Modular package structure
- Interface-based design
- Separation of concerns
- Clean error handling
- Context-aware HTTP middleware

---

## License
This project is open source and available under the [MIT License](LICENSE).
