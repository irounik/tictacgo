package repository

import (
	"fmt"
	"tictacgo/server/db"
	"tictacgo/server/model/entity"
)

type ServerGameRepo interface {
	Save(game *entity.ServerGame) error
	GetById(id string) (*entity.ServerGame, error)
}

type InMemoryGameRepo struct {
	db *db.InMemoryDb[string, *entity.ServerGame]
}

func NewGameRepository() *InMemoryGameRepo {
	return &InMemoryGameRepo{db: db.NewInMemoryDb[string, *entity.ServerGame]()}
}

func (repo *InMemoryGameRepo) Save(game *entity.ServerGame) error {
	repo.db.Save(game.Id, game)
	return nil
}

func (repo *InMemoryGameRepo) GetById(gameId string) (*entity.ServerGame, error) {
	game, exists := repo.db.Get(gameId)
	if !exists {
		return nil, fmt.Errorf("game not found")
	}
	return game, nil
}
