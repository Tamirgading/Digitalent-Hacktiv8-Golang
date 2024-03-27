package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"unique;not null" binding:"required"`
	Email     string         `json:"email" gorm:"unique;not null" binding:"required;email"`
	Password  string         `json:"password" gorm:"not null" binding:"required"`
	Age       uint8          `json:"age" binding:"required,gte=8"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
