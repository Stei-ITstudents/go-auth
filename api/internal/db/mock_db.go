package db

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type InMemoryUserStore struct {
	users map[string]string // username -> hashed password
}

func NewInMemoryUserStore() (*InMemoryUserStore, error) {
	return &InMemoryUserStore{
		users: make(map[string]string),
	}, nil
}

func (s *InMemoryUserStore) Get(username string) (string, error) {
	hashedPassword, exists := s.users[username]
	if !exists {
		return "", errors.New("user not found")
	}

	return hashedPassword, nil
}

func (s *InMemoryUserStore) Add(username, password string) error {
	if _, exists := s.users[username]; exists {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.users[username] = string(hashedPassword)
	return nil
}
