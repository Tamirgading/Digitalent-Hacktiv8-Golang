package model

import "gorm.io/gorm"

type SocialMedia struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name" binding:"required"`
	SocialMediaURL  string         `json:"social_media_url" binding:"required"`
	UserID          uint           `json:"user_id" binding:"required"`
	CreatedAt       *gorm.DeletedAt `json:"created_at"`
	UpdatedAt       *gorm.DeletedAt `json:"updated_at"`
	User            User           `json:"User" gorm:"foreignKey:UserID"`
}
