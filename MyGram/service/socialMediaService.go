package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
)

type SocialMedia struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uuid.UUID `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaService struct {
	repo *SocialMediaRepository
}

func NewSocialMediaService(repo *SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{repo: repo}
}

func (s *SocialMediaService) CreateSocialMedia(socialMedia *SocialMedia) (*SocialMedia, error) {
	socialMedia.ID = uuid.New()
	socialMedia.CreatedAt = time.Now()
	socialMedia.UpdatedAt = time.Now()

	err := s.repo.CreateSocialMedia(socialMedia)
	if err != nil {
		log.Printf("Failed to create social media: %v", err)
		return nil, err
	}
	return socialMedia, nil
}

func (s *SocialMediaService) GetSocialMedias() ([]*SocialMedia, error) {
	socialMedias, err := s.repo.GetSocialMedias()
	if err != nil {
		log.Printf("Failed to fetch social medias: %v", err)
		return nil, err
	}
	return socialMedias, nil
}

func (s *SocialMediaService) UpdateSocialMedia(socialMediaID uuid.UUID, updatedSocialMedia *SocialMedia) (*SocialMedia, error) {
	updatedSocialMedia.UpdatedAt = time.Now()
	err := s.repo.UpdateSocialMedia(socialMediaID, updatedSocialMedia)
	if err != nil {
		log.Printf("Failed to update social media: %v", err)
		return nil, err
	}
	return updatedSocialMedia, nil
}

func (s *SocialMediaService) DeleteSocialMedia(socialMediaID uuid.UUID) error {
	err := s.repo.DeleteSocialMedia(socialMediaID)
	if err != nil {
		log.Printf("Failed to delete social media: %v", err)
		return err
	}
	return nil
}
