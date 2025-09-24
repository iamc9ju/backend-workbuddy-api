package service

import (
	"app/src/enum"
	"app/src/model"
	"app/src/repository"
	"app/src/utils"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type JobService interface {
	GetProjectList(ctx context.Context) ([]model.Project, error)
	CreateProject(ctx context.Context, body model.ProjectCreate, ownderID uint) (*model.Project, error)
	GetProjectByProjectID(ctx context.Context, projectID uint) (*model.Project, error)
	GetProjectBySlug(ctx context.Context, slug string) (*model.Project, error)
	GetProjectsByOwnerID(ctx context.Context, ownerID uint) ([]model.Project, error)
}
type jobService struct {
	repo     repository.JobRepository
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewJobService(repo repository.JobRepository, validate *validator.Validate) JobService {
	return &jobService{
		repo:     repo,
		Log:      utils.Log,
		Validate: validate,
	}
}

func (s *jobService) GetProjectList(ctx context.Context) ([]model.Project, error) {
	s.Log.Info("Retrieving project list")
	projects, err := s.repo.GetProjectList(ctx)
	if err != nil {
		s.Log.WithError(err).Error("Failed to retrieve project list")
		return nil, err
	}
	s.Log.Infof("Successfully retrieved %d projects", len(projects))
	return projects, nil
}

func (s *jobService) CreateProject(ctx context.Context, body model.ProjectCreate, ownerID uint) (*model.Project, error) {
	if err := s.Validate.Struct(body); err != nil {
		s.Log.WithError(err).Error("Validation failed")
		return nil, err
	}

	project := &model.Project{
		OwnerID:     body.OwnerID,
		Title:       body.Title,
		Description: body.Description,
		BudgetMin:   body.BudgetMin,
		BudgetMax:   body.BudgetMax,
		Currency:    body.Currency,
		Deadline:    body.Deadline,
		Status:      enum.DRAFT_PROJECT, // หรือ "open" ตาม business logic
	}

	if err := s.repo.CreateProject(ctx, project); err != nil {
		s.Log.WithError(err).Error("Failed to create project")
		return nil, err
	}

	return project, nil

}

func (s *jobService) GetProjectByProjectID(ctx context.Context, projectID uint) (*model.Project, error) {
	project, err := s.repo.GetProjectByProjectID(ctx, projectID)
	if err != nil {
		s.Log.WithError(err).Errorf("Failed to get project with project ID: %d", projectID)
		return nil, err
	}
	return project, nil
}

func (s *jobService) GetProjectBySlug(ctx context.Context, slug string) (*model.Project, error) {
	project, err := s.repo.GetProjectBySlug(ctx, slug)
	if err != nil {
		s.Log.WithError(err).Errorf("Failed to get project with project slug: %d", slug)
		return nil, err
	}
	return project, nil
}

func (s *jobService) GetProjectsByOwnerID(ctx context.Context, ownerID uint) ([]model.Project, error) {
	projects, err := s.repo.GetProjectsByOwnerID(ctx, ownerID)
	if err != nil {
		s.Log.WithError(err).Errorf("Failed to get projects for owner ID: %d", ownerID)
		return nil, err
	}
	return projects, nil
}
