package service

import (
	"app/src/model"
	"app/src/repository"
	"app/src/utils"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type CategoryService interface {
	GetCategoryList(ctx context.Context) ([]model.Category, error)
	CreateCategory(ctx context.Context, body model.CategoryCreate) (*model.Category, error)
	GetCategoryByCategoryID(ctx context.Context, categoryID uint) (*model.Category, error)
}
type categoryService struct {
	repo     repository.CategoryRepository
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewCategoryService(repo repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &categoryService{
		repo:     repo,
		Log:      utils.Log,
		Validate: validate,
	}
}

func (s *categoryService) GetCategoryList(ctx context.Context) ([]model.Category, error) {
	s.Log.Info("Getting category list")
	category, err := s.repo.GetCategoryList(ctx)
	if err != nil {
		s.Log.WithError(err).Error("Failed to get category")
		return nil, err
	}
	s.Log.Infof("Got %d category", len(category))
	return category, nil
}

func (s *categoryService) CreateCategory(ctx context.Context, body model.CategoryCreate) (*model.Category, error) {
	if err := s.Validate.Struct(body); err != nil {
		s.Log.WithError(err).Error("Validation failed")
		return nil, err
	}

	category := &model.Category{
		Title:       body.Title,
		Description: body.Description,
	}

	if err := s.repo.CreateCategory(ctx, category); err != nil {
		s.Log.WithError(err).Error("Failed to create category")
		return nil, err
	}

	return category, nil

}

func (s *categoryService) GetCategoryByCategoryID(ctx context.Context, categoryID uint) (*model.Category, error) {
	category, err := s.repo.GetCategoryByCategoryID(ctx, categoryID)
	if err != nil {
		s.Log.WithError(err).Errorf("Failed to get Category with category ID: %d", categoryID)
		return nil, err
	}
	return category, nil
}
