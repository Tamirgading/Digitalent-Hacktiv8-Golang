package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

type AuthService struct {
	db *pgx.Conn
}

func NewAuthService(db *pgx.Conn) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) StoreToken(token string, expiration time.Time) error {
	_, err := s.db.Exec(context.Background(), "INSERT INTO tokens (token, expiration) VALUES ($1, $2)", token, expiration)
	if err != nil {
		log.Printf("Failed to store token in database: %v", err)
		return err
	}
	return nil
}

func (s *AuthService) DeleteExpiredTokens() error {
	_, err := s.db.Exec(context.Background(), "DELETE FROM tokens WHERE expiration < now()")
	if err != nil {
		log.Printf("Failed to delete expired tokens from database: %v", err)
		return err
	}
	return nil
}
