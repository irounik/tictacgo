package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	db "tictacgo/server/db"
	"time"
)

type Session struct {
	Username string
	Token    string
	ExpireAt time.Time
}

type SessionService interface {
	Create(username string) (*Session, error)
	GetByToken(token string) (*Session, error)
	Invalidate(token string) error
}

type InMemorySessionService struct {
	sessionByToken    *db.InMemoryDb[string, *Session]
	sessionByUsername *db.InMemoryDb[string, *Session]
}

func NewSessionService() *InMemorySessionService {
	return &InMemorySessionService{
		sessionByToken:    db.NewInMemoryDb[string, *Session](),
		sessionByUsername: db.NewInMemoryDb[string, *Session](),
	}
}

func (s *InMemorySessionService) Create(username string) (*Session, error) {
	token, err := s.generateToken()
	if err != nil {
		return nil, err
	}

	newSession := &Session{
		Username: username,
		Token:    token,
		ExpireAt: time.Now().UTC().Add(1 * time.Hour),
	}
	s.sessionByToken.Save(token, newSession)
	s.sessionByUsername.Save(username, newSession)
	return newSession, nil
}

func (s *InMemorySessionService) GetByToken(token string) (*Session, error) {
	session, exists := s.sessionByToken.Get(token)
	if !exists {
		return nil, fmt.Errorf("session not found")
	}

	if session.ExpireAt.Before(time.Now().UTC()) {
		return nil, fmt.Errorf("session expired")
	}

	return session, nil
}

func (s *InMemorySessionService) Invalidate(token string) error {
	session, exists := s.sessionByToken.Get(token)
	if !exists {
		return nil
	}

	s.sessionByToken.Delete(token)
	s.sessionByUsername.Delete(session.Username)
	return nil
}

func (s *InMemorySessionService) generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to read random bytes: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
