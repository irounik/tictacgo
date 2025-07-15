package main

import (
	"flag"
	"fmt"
	"os"
	"tictacgo/cli"
)

func main() {
	mode := flag.String("mode", "cli", "Specify game mode, options are: ['cli', 'server']")
	flag.Parse()

	// CLI Game
	if len(os.Args) < 2 || mode == nil || *mode == "cli" {
		game := cli.SetupGame()
		cli.Play(game)
		return
	}

	if *mode == "server" {
		fmt.Println("HTTP mode is not implemented yet!")
		return
	}

	fmt.Printf("'%s' mode is not supported\n", *mode)
}
