package service

import (
	"app/src/model"
	"app/src/repository"
	"app/src/utils"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetAllUser(ctx context.Context) ([]model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewUserService(userRepo repository.UserRepository, validate *validator.Validate) UserService {
	return &userService{
		userRepo: userRepo,
		Log:      utils.Log,
		Validate: validate,
	}
}

func (us *userService) GetAllUser(ctx context.Context) ([]model.User, error) {
	return us.userRepo.GetAllUser(ctx)
}
