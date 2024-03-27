package repository

import (
	"project-pertama/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) Create(user *model.User) error {
	result := ur.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	result := ur.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) GetByID(userID uint) (*model.User, error) {
	var user model.User
	result := ur.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) Update(user *model.User) error {
	result := ur.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) Delete(userID uint) error {
	result := ur.DB.Delete(&model.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
