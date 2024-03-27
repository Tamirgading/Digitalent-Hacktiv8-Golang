package model

import "gorm.io/gorm"

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Message   string         `json:"message" binding:"required"`
	UserID    uint           `json:"user_id" binding:"required"`
	PhotoID   uint           `json:"photo_id" binding:"required"`
	CreatedAt *gorm.DeletedAt `json:"created_at"`
	UpdatedAt *gorm.DeletedAt `json:"updated_at"`
	User      User           `json:"User" gorm:"foreignKey:UserID"`
	Photo     Photo          `json:"Photo" gorm:"foreignKey:PhotoID"`
}
