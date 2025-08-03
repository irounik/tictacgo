package cli

import (
	"fmt"
	"tictacgo/game"
)

func SetupGame() game.Game {
	var len int = 3
	rows := make([][]string, len)
	for i := 0; i < len; i++ {
		rows[i] = make([]string, len)
		for j := 0; j < len; j++ {
			rows[i][j] = ""
		}
	}

	var board game.Board = game.Board{Rows: rows, Size: len}
	var firstPlayerName, secondPlayerName string
	fmt.Printf("Enter name for first payler (X): ")
	fmt.Scanln(&firstPlayerName)

	fmt.Printf("Enter name for second payler (O): ")
	fmt.Scanln(&secondPlayerName)

	var player1 = game.Player{Name: firstPlayerName, Symbol: "X"}
	var player2 = game.Player{Name: secondPlayerName, Symbol: "O"}

	return game.Game{Board: board, Players: []*game.Player{&player1, &player2}, CurrentTurn: &player1}
}

func Play(g game.Game) {
	for {
		g.Board.Print()
		player := g.CurrentTurn

		fmt.Printf("%s (%s), enter row and column: ", player.Name, player.Symbol)
		var row, col int
		fmt.Scanln(&row, &col)

		err := g.Board.Mark(row, col, player.Symbol)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if g.Board.IsWinningMove(row, col, player.Symbol) {
			fmt.Printf("%s (%s) wins!\n", player.Name, player.Symbol)
			g.Board.Print()
			break
		}

		if g.Board.IsFull() {
			fmt.Println("Board is full, game over")
			break
		}

		g.UpdateTurn()
	}
}
