package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoService struct {
	repo *PhotoRepository
}

func NewPhotoService(repo *PhotoRepository) *PhotoService {
	return &PhotoService{repo: repo}
}

func (s *PhotoService) CreatePhoto(photo *Photo) (*Photo, error) {
	photo.ID = uuid.New()
	photo.CreatedAt = time.Now()
	photo.UpdatedAt = time.Now()

	err := s.repo.CreatePhoto(photo)
	if err != nil {
		log.Printf("Failed to create photo: %v", err)
		return nil, err
	}
	return photo, nil
}

func (s *PhotoService) GetPhotos() ([]*Photo, error) {
	photos, err := s.repo.GetPhotos()
	if err != nil {
		log.Printf("Failed to fetch photos: %v", err)
		return nil, err
	}
	return photos, nil
}

func (s *PhotoService) UpdatePhoto(photoID uuid.UUID, updatedPhoto *Photo) (*Photo, error) {
	updatedPhoto.UpdatedAt = time.Now()
	err := s.repo.UpdatePhoto(photoID, updatedPhoto)
	if err != nil {
		log.Printf("Failed to update photo: %v", err)
		return nil, err
	}
	return updatedPhoto, nil
}

func (s *PhotoService) DeletePhoto(photoID uuid.UUID) error {
	err := s.repo.DeletePhoto(photoID)
	if err != nil {
		log.Printf("Failed to delete photo: %v", err)
		return err
	}
	return nil
}
