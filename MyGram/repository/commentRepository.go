package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(comment *Comment) error {
	query := `INSERT INTO comments (id, user_id, photo_id, message, created_at, updated_at)
			   VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, comment.ID, comment.UserID, comment.PhotoID, comment.Message, comment.CreatedAt, comment.UpdatedAt)
	if err != nil {
		log.Printf("Failed to insert comment into database: %v", err)
		return err
	}
	return nil
}

func (r *CommentRepository) GetComments() ([]*Comment, error) {
	query := `SELECT * FROM comments`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Failed to get comments from database: %v", err)
		return nil, err
	}
	defer rows.Close()

	comments := []*Comment{}
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			log.Printf("Failed to scan comment row: %v", err)
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over comment rows: %v", err)
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) UpdateComment(commentID uuid.UUID, updatedComment *Comment) error {
	query := `UPDATE comments 
			  SET message = $1, updated_at = $2 
			  WHERE id = $3`
	_, err := r.db.Exec(query, updatedComment.Message, updatedComment.UpdatedAt, commentID)
	if err != nil {
		log.Printf("Failed to update comment in database: %v", err)
		return err
	}
	return nil
}

func (r *CommentRepository) DeleteComment(commentID uuid.UUID) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.db.Exec(query, commentID)
	if err != nil {
		log.Printf("Failed to delete comment from database: %v", err)
		return err
	}
	return nil
}
