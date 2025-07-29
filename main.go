package main

import (
	"flag"
	"fmt"
	"os"
	"tictacgo/cli"
	"tictacgo/server"
)

func main() {
	mode := flag.String("mode", "cli", "Specify game mode, options are: ['cli', 'server']")
	port := flag.Int("port", 8080, "Port for server to start on")
	flag.Parse()

	// CLI Game
	if len(os.Args) < 2 || mode == nil || *mode == "cli" {
		game := cli.SetupGame()
		cli.Play(game)
		return
	}

	// HTTP Server
	if *mode == "server" {
		server.Start(*port)
		return
	}

	fmt.Printf("'%s' mode is not supported\n", *mode)
}
