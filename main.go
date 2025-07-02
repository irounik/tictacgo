package main

import (
	"tictacgo/game"
)

func main() {
	var len int = 3
	rows := make([][]string, len)
	for i := 0; i < len; i++ {
		rows[i] = make([]string, len)
		for j := 0; j < len; j++ {
			rows[i][j] = ""
		}
	}

	var board game.Board = game.Board{Rows: rows, Size: len}
	var player1 = game.Player{Name: "Rounik", Symbol: "X"}
	var player2 = game.Player{Name: "Rohan", Symbol: "O"}

	var game = game.Game{Board: board, Players: []game.Player{player1, player2}}
	game.Play()
}
