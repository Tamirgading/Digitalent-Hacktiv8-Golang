package repository

import (
	"project-pertama/model"

	"gorm.io/gorm"
)

type PhotoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		DB: db,
	}
}

func (pr *PhotoRepository) Create(photo *model.Photo) error {
	result := pr.DB.Create(photo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *PhotoRepository) GetByID(photoID uint) (*model.Photo, error) {
	var photo model.Photo
	result := pr.DB.First(&photo, photoID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &photo, nil
}

func (pr *PhotoRepository) Update(photo *model.Photo) error {
	result := pr.DB.Save(photo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *PhotoRepository) Delete(photoID uint) error {
	result := pr.DB.Delete(&model.Photo{}, photoID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
