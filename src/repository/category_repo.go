package repository

import (
	"app/src/model"
	"errors"
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoryList(ctx context.Context) ([]model.Category, error)
	CreateCategory(ctx context.Context, category *model.Category) error
	GetCategoryByCategoryID(ctx context.Context, categoryID uint) (*model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetCategoryList(ctx context.Context) ([]model.Category, error) {
	var category []model.Category

	err := r.db.WithContext(ctx).Find(&category).Error

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) CreateCategory(ctx context.Context, category *model.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryRepository) GetCategoryByCategoryID(ctx context.Context, categoryID uint) (*model.Category, error) {
	var category model.Category
	err := r.db.WithContext(ctx).
		Where("category_id = ?", categoryID).
		First(&category).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("categorynot found or you don't have permission")
		}
		return nil, err
	}
	return &category, nil
}
