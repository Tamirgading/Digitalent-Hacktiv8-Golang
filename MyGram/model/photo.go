package model

import "gorm.io/gorm"

type Photo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" binding:"required"`
	Caption   string         `json:"caption"`
	PhotoURL  string         `json:"photo_url" binding:"required"`
	UserID    uint           `json:"user_id" binding:"required"`
	CreatedAt *gorm.DeletedAt `json:"created_at"`
	UpdatedAt *gorm.DeletedAt `json:"updated_at"`
	User      User           `json:"User" gorm:"foreignKey:UserID"`
}
