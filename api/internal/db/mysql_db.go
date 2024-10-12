package db

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type MySqlUserStore struct {
	db *sql.DB
}

func NewMySqlUserStore(dns string) (*MySqlUserStore, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &MySqlUserStore{db: db}, nil
}

func (s *MySqlUserStore) Get(username string) (string, error) {
	var hashedPassword string
	err := s.db.QueryRow("SELECT password_hash FROM users WHERE username=?", username).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}

	return hashedPassword, nil

}

func (s *MySqlUserStore) Add(username, password string) error {
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=?", username).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.db.Exec("INSERT INTO users (username, password_hash) VALUES(?,?)", username, hashedPassword)

	return err

}
