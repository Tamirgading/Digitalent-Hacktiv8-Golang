package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uuid.UUID `json:"photo_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentService struct {
	repo *CommentRepository
}

func NewCommentService(repo *CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment *Comment) (*Comment, error) {
	comment.ID = uuid.New()
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	err := s.repo.CreateComment(comment)
	if err != nil {
		log.Printf("Failed to create comment: %v", err)
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) GetComments() ([]*Comment, error) {
	comments, err := s.repo.GetComments()
	if err != nil {
		log.Printf("Failed to fetch comments: %v", err)
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) UpdateComment(commentID uuid.UUID, updatedComment *Comment) (*Comment, error) {
	updatedComment.UpdatedAt = time.Now()
	err := s.repo.UpdateComment(commentID, updatedComment)
	if err != nil {
		log.Printf("Failed to update comment: %v", err)
		return nil, err
	}
	return updatedComment, nil
}

func (s *CommentService) DeleteComment(commentID uuid.UUID) error {
	err := s.repo.DeleteComment(commentID)
	if err != nil {
		log.Printf("Failed to delete comment: %v", err)
		return err
	}
	return nil
}
