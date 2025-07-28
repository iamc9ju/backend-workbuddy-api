package repository

import (
	"app/src/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUser(ctx context.Context) ([]model.User, error) {
	var users []model.User

	err := r.db.WithContext(ctx).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}
