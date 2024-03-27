package repository

import (
	"project-pertama/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		DB: db,
	}
}

func (cr *CommentRepository) Create(comment *model.Comment) error {
	result := cr.DB.Create(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CommentRepository) GetByID(commentID uint) (*model.Comment, error) {
	var comment model.Comment
	result := cr.DB.First(&comment, commentID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (cr *CommentRepository) Update(comment *model.Comment) error {
	result := cr.DB.Save(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CommentRepository) Delete(commentID uint) error {
	result := cr.DB.Delete(&model.Comment{}, commentID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
