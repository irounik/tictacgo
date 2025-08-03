package service

import (
	"fmt"
	"tictacgo/game"
	"tictacgo/server/model/entity"
	"tictacgo/server/repository"

	"github.com/google/uuid"
)

type GameService interface {
	NewGame(hostUsername string) (*entity.ServerGame, error)
	JoinGame(gameId string, guestUsername string) error
	StartGame(gameId string, hostUsername string) error
	GetGame(gameId string, player string) (*entity.ServerGame, error)
	MakeMove(gameId string, player string, row int, col int) (*entity.ServerGame, error)
	GetSymbol(game *entity.ServerGame, player string) string
}

type ServerGameService struct {
	gameRepo repository.ServerGameRepo
}

func NewGameService(gameRepo repository.ServerGameRepo) *ServerGameService {
	return &ServerGameService{gameRepo: gameRepo}
}

func (s *ServerGameService) NewGame(hostUsername string) (*entity.ServerGame, error) {

	len := 3
	rows := make([][]string, len)
	for i := 0; i < len; i++ {
		rows[i] = make([]string, len)
		for j := 0; j < len; j++ {
			rows[i][j] = ""
		}
	}

	board := game.Board{
		Rows: rows,
		Size: len,
	}

	players := make([]*game.Player, 2)
	players[0] = &game.Player{
		Name:   hostUsername,
		Symbol: "X",
	}

	players[1] = nil // Guest player will join later

	game := &game.Game{
		Board:       board,
		Players:     players,
		CurrentTurn: players[0],
	}

	newGame := &entity.ServerGame{
		Id:             uuid.NewString(),
		Game:           game,
		HostUsername:   &hostUsername,
		GuestUsername:  nil,
		WinnerUsername: nil,
		Status:         entity.New,
	}

	err := s.gameRepo.Save(newGame)
	if err != nil {
		return nil, err
	}

	return newGame, nil
}

func (s *ServerGameService) JoinGame(gameId string, guestUsername string) error {
	g, err := s.gameRepo.GetById(gameId)
	if err != nil {
		return fmt.Errorf("game not found")
	}

	if g.GuestUsername != nil {
		return fmt.Errorf("game is full")
	}

	if *g.HostUsername == guestUsername {
		return fmt.Errorf("can't join your own game as guest")
	}

	g.GuestUsername = &guestUsername
	g.Game.Players[1] = &game.Player{
		Name:   guestUsername,
		Symbol: "O",
	}
	return nil
}

func (s *ServerGameService) StartGame(gameId string, hostUsername string) error {
	game, err := s.gameRepo.GetById(gameId)
	if err != nil {
		return err
	}

	if *game.HostUsername != hostUsername {
		return fmt.Errorf("game not found")
	}

	if *game.GuestUsername == hostUsername {
		return fmt.Errorf("only hosts can start a game")
	}

	if game.GuestUsername == nil {
		return fmt.Errorf("guest has not joined the game yet")
	}

	if game.Status != entity.New {
		return fmt.Errorf("cant start a '%d' game", game.Status)
	}

	game.Status = entity.Ongoing
	return nil
}

func (s *ServerGameService) GetGame(gameId string, player string) (*entity.ServerGame, error) {
	game, err := s.gameRepo.GetById(gameId)
	if err != nil {
		return nil, err
	}

	if (player != *game.HostUsername) && (game.GuestUsername != nil && *game.GuestUsername != player) {
		return nil, fmt.Errorf("game not found")
	}

	return game, nil
}

func (s *ServerGameService) MakeMove(gameId string, player string, row int, col int) (*entity.ServerGame, error) {
	game, err := s.GetGame(gameId, player)
	if err != nil {
		return nil, err
	}

	if player != game.Game.CurrentTurn.Name {
		return nil, fmt.Errorf("not your turn")
	}

	symbol := s.GetSymbol(game, player)
	err = game.Game.Board.Mark(row, col, symbol)
	if err != nil {
		return nil, fmt.Errorf("invalid move: %s", err.Error())
	}

	if game.Game.Board.IsWinningMove(row, col, symbol) {
		game.Status = entity.Complete
	}

	game.Game.UpdateTurn()
	return game, nil
}

func (s *ServerGameService) GetSymbol(game *entity.ServerGame, player string) string {
	if *game.HostUsername == player {
		return "X"
	}

	return "O"
}
