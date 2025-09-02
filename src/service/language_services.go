package service

import (
	"app/src/model"
	"app/src/repository"
	"app/src/utils"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type LanguageService interface {
	ListAllLanguages() ([]model.ProgrammingLanguage, error)
}

type languageService struct {
	rp       repository.ProgrammingLanguageRepository
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewLanguageService(rp repository.ProgrammingLanguageRepository, validate *validator.Validate) LanguageService {
	return &languageService{
		rp:       rp,
		Log:      utils.Log,
		Validate: validate,
	}
}

func (s *languageService) ListAllLanguages() ([]model.ProgrammingLanguage, error) {
	languages, err := s.rp.ListAllLanguages()
	if err != nil {
		s.Log.WithError(err).Error("failed to retrieve programming languages from repository")
		return nil, err
	}
	s.Log.WithFields(logrus.Fields{
		"count": len(languages),
	}).Info("successfully retrieved all programming languages")

	return languages, nil
}
