package game

import "fmt"

type Game struct {
	Board   Board
	Players []Player
}

func (g *Game) Play() {
	currentPlayer := 0
	for {
		g.Board.Print()
		player := g.Players[currentPlayer]

		fmt.Printf("%s (%s), enter row and column: ", player.Name, player.Symbol)
		var row, col int
		fmt.Scanln(&row, &col)

		err := g.Board.Mark(row, col, player.Symbol)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if g.Board.IsFull() {
			fmt.Println("Board is full, game over")
			break
		}

		if g.Board.IsWinningMove(row, col, player.Symbol) {
			fmt.Printf("%s (%s) wins!\n", player.Name, player.Symbol)
			g.Board.Print()
			break
		}

		currentPlayer = (currentPlayer + 1) % len(g.Players)
	}
}
