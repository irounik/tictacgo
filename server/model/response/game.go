package response

type GameResponse struct {
	GameId      string     `json:"gameId"`
	CurrentTurn string     `json:"currentTurn"`
	YourSymbol  string     `json:"yourSymbol"`
	Board       [][]string `json:"board"`
}
