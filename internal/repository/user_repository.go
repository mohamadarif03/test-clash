package repository

import (
	"test-clash-be/internal/model"
	"test-clash-be/pkg/database"
)

func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

func FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
