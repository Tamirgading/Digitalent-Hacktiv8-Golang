package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *User) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, err
	}
	user.Password = string(hashedPassword)
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = s.repo.CreateUser(user)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("Invalid password: %v", err)
		return "", err
	}

	// Generate JWT token
	token, err := GenerateJWTToken(user.ID)
	if err != nil {
		log.Printf("Failed to generate JWT token: %v", err)
		return "", err
	}
	return token, nil
}

func (s *UserService) UpdateUser(userID uuid.UUID, updatedUser *User) (*User, error) {
	updatedUser.UpdatedAt = time.Now()
	err := s.repo.UpdateUser(userID, updatedUser)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return nil, err
	}
	return updatedUser, nil
}

func (s *UserService) DeleteUser(userID uuid.UUID) error {
	err := s.repo.DeleteUser(userID)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return err
	}
	return nil
}
