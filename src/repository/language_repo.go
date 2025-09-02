package repository

import (
	"app/src/model"
	"time"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProgrammingLanguageRepository interface {
	ListAllLanguages() ([]model.ProgrammingLanguage, error)
}

type programmingLanguageRepository struct {
	db *gorm.DB
}

func NewProgrammingLanguageRepository(db *gorm.DB) ProgrammingLanguageRepository {
	return &programmingLanguageRepository{db: db}
}

func (r *programmingLanguageRepository) ListAllLanguages() ([]model.ProgrammingLanguage, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	var languages []model.ProgrammingLanguage

	result := r.db.WithContext(ctx).Select("language_id", "language_name").Order("language_name ASC").Find(&languages)

	if result.Error != nil {
		return nil, result.Error
	}

	return languages, nil
}
