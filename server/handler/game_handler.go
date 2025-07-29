package handler

import (
	"net/http"
)

type GameHandler interface {
	NewGame(w http.ResponseWriter, r *http.Request)
}

type GameHandlerImpl struct {
}

func (gh *GameHandlerImpl) NewGame(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this
	w.Write([]byte("TODO: Implement this!"))
}
