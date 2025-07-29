package service

import (
	"fmt"
	"tictacgo/server/model/entity"
	"tictacgo/server/model/request"
	"tictacgo/server/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.UserRepo
}

func NewAuthService(userRepo repository.UserRepo) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) SignUp(authRequest request.AuthRequest) error {

	userExists, err := s.userRepo.CheckExists(authRequest.Username)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}

	if userExists {
		return fmt.Errorf("user with username '%s' already exists", authRequest.Username)
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(authRequest.Password), 16)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}

	hashedPassword := string(hashedBytes)
	newUser := &entity.User{Username: authRequest.Username, HashedPassword: hashedPassword}
	if err := s.userRepo.Save(*newUser); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func (s AuthService) SignIn(authRequest request.AuthRequest) error {
	user, err := s.userRepo.GetById(authRequest.Username)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}

	hashedPassword := []byte(user.HashedPassword)
	dbPassword := []byte(authRequest.Password)
	if err := bcrypt.CompareHashAndPassword(hashedPassword, dbPassword); err != nil {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}
