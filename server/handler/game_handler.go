package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tictacgo/server/config"
	"tictacgo/server/model/request"
	"tictacgo/server/model/response"
	"tictacgo/server/service"
)

type GameHandler interface {
	NewGame(w http.ResponseWriter, r *http.Request)
	JoinGame(w http.ResponseWriter, r *http.Request)
	StartGame(w http.ResponseWriter, r *http.Request)
	GetGame(w http.ResponseWriter, r *http.Request)
	MakeMove(w http.ResponseWriter, r *http.Request)
}

type GameHandlerImpl struct {
	gameService service.GameService
}

func NewGameHandler(gameService service.GameService) *GameHandlerImpl {
	return &GameHandlerImpl{gameService: gameService}
}

func (gh *GameHandlerImpl) NewGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
		return
	}

	hostUsername, ok := r.Context().Value(config.AUTH_USERNAME).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid username!"))
	}

	newGame, err := gh.gameService.NewGame(hostUsername)
	if err != nil {
		log.Println("error occoured", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create game"))
	}

	log.Printf("Created new game with ID: %s", newGame.Id)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(newGame.Id))
}

func (gh *GameHandlerImpl) JoinGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
		return
	}

	guestUsername, ok := r.Context().Value(config.AUTH_USERNAME).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid username!"))
		return
	}

	gameID := r.PathValue("gameId")
	if gameID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid game id!"))
		return
	}

	gh.gameService.JoinGame(gameID, guestUsername)
}

func (gh *GameHandlerImpl) StartGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
		return
	}

	hostUsername, ok := r.Context().Value(config.AUTH_USERNAME).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid username!"))
		return
	}

	gameId := r.PathValue("gameId")
	if gameId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid game id!"))
		return
	}

	err := gh.gameService.StartGame(gameId, hostUsername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Started game!"))
}

func (gh *GameHandlerImpl) GetGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
		return
	}

	requester, ok := r.Context().Value(config.AUTH_USERNAME).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid username!"))
		return
	}

	gameId := r.PathValue("gameId")
	if gameId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid game id!"))
		return
	}

	game, err := gh.gameService.GetGame(gameId, requester)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	symbol := gh.gameService.GetSymbol(game, requester)
	response := response.GameResponse{
		GameId:      gameId,
		CurrentTurn: game.Game.CurrentTurn.Name,
		YourSymbol:  symbol,
		Board:       game.Game.Board.Rows,
	}

	respJson, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func (gh *GameHandlerImpl) MakeMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
		return
	}

	player, ok := r.Context().Value(config.AUTH_USERNAME).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid username!"))
		return
	}

	gameId := r.PathValue("gameId")
	if gameId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid game id!"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var moveRequest request.MoveRequest
	decoder.Decode(&moveRequest)

	upatedGame, err := gh.gameService.MakeMove(gameId, player, moveRequest.Row, moveRequest.Col)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	symbol := gh.gameService.GetSymbol(upatedGame, player)
	response := response.GameResponse{
		GameId:      gameId,
		CurrentTurn: upatedGame.Game.CurrentTurn.Name,
		YourSymbol:  symbol,
		Board:       upatedGame.Game.Board.Rows,
	}

	respJson, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respJson))
}
