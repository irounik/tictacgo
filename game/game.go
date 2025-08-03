package game

type Game struct {
	Board       Board
	Players     []*Player
	CurrentTurn *Player
}

func (g *Game) UpdateTurn() {
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i] != g.CurrentTurn {
			continue
		}

		// Found current player index
		playerIdx := (i + 1) % len(g.Players)
		g.CurrentTurn = g.Players[playerIdx]
		break
	}
}
