package repository

import (
	"project-pertama/model"

	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		DB: db,
	}
}

func (smr *SocialMediaRepository) Create(socialMedia *model.SocialMedia) error {
	result := smr.DB.Create(socialMedia)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (smr *SocialMediaRepository) GetByID(socialMediaID uint) (*model.SocialMedia, error) {
	var socialMedia model.SocialMedia
	result := smr.DB.First(&socialMedia, socialMediaID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &socialMedia, nil
}

func (smr *SocialMediaRepository) Update(socialMedia *model.SocialMedia) error {
	result := smr.DB.Save(socialMedia)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (smr *SocialMediaRepository) Delete(socialMediaID uint) error {
	result := smr.DB.Delete(&model.SocialMedia{}, socialMediaID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
