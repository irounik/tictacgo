package entity

import (
	"tictacgo/game"
)

type Status int

const (
	New Status = iota
	Ongoing
	Complete
)

type ServerGame struct {
	Id             string
	Game           *game.Game
	HostUsername   *string
	GuestUsername  *string
	WinnerUsername *string
	Status         Status
}
