package repository

import (
	"fmt"
	"tictacgo/server/db"
	"tictacgo/server/model/entity"
)

type UserRepo interface {
	Save(entity.User) error
	GetById(username string) (*entity.User, error)
	CheckExists(username string) (bool, error)
	UpdatePassword(username string, hashedPassword string) error
	Delete(username string)
}

type InMemoryUserRepo struct {
	db *db.InMemoryDb[string, entity.User]
}

func NewUserRepository() *InMemoryUserRepo {
	return &InMemoryUserRepo{db: db.NewInMemoryDb[string, entity.User]()}
}

func (repo *InMemoryUserRepo) Save(user entity.User) error {
	repo.db.Save(user.Username, user)
	return nil
}

func (repo *InMemoryUserRepo) GetById(username string) (*entity.User, error) {
	user, exists := repo.db.Get(username)
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (repo *InMemoryUserRepo) CheckExists(username string) (bool, error) {
	exists := repo.db.Exists(username)
	return exists, nil
}

func (repo *InMemoryUserRepo) UpdatePassword(username string, hashedPassword string) error {
	user, exists := repo.db.Get(username)
	if !exists {
		return fmt.Errorf("user not found")
	}
	user.HashedPassword = hashedPassword
	repo.db.Save(username, user)
	return nil
}

func (repo *InMemoryUserRepo) Delete(username string) {
	repo.db.Delete(username)
}
