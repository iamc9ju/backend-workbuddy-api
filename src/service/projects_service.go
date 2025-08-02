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

type ProjectService interface {
	GetProjectList(ctx context.Context) ([]model.Project, error)
	CreateProject(ctx context.Context, input model.ProjectCreateInput, ownderID uint) (*model.Project, error)
	GetProjectByProjectID(ctx context.Context, projectID uint) (*model.Project, error)
	GetProjectsByOwnerID(ctx context.Context, ownerID uint) ([]model.Project, error)
}
type projectService struct {
	repo     repository.ProjectRepository
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewProjectService(repo repository.ProjectRepository, validate *validator.Validate) ProjectService {
	return &projectService{
		repo:     repo,
		Log:      utils.Log,
		Validate: validate,
	}
}

func (s *projectService) GetProjectList(ctx context.Context) ([]model.Project, error) {
	s.Log.Info("Getting project list")
	projects, err := s.repo.GetProjectList(ctx)
	if err != nil {
		s.Log.WithError(err).Error("Failed to get projects")
		return nil, err
	}
	s.Log.Infof("Got %d projects", len(projects))
	return projects, nil
}

func (s *projectService) CreateProject(ctx context.Context, input model.ProjectCreateInput, ownerID uint) (*model.Project, error) {
	if err := s.Validate.Struct(input); err != nil {
		s.Log.WithError(err).Error("Validation failed")
		return nil, err
	}

	project := &model.Project{
		OwnerID:     input.OwnerID,
		Title:       input.Title,
		Description: input.Description,
		Budget:      input.Budget,
		Currency:    input.Currency,
		Deadline:    input.Deadline,
		Status:      enum.DRAFT_PROJECT, // หรือ "open" ตาม business logic
	}

	if err := s.repo.CreateProject(ctx, project); err != nil {
		s.Log.WithError(err).Error("Failed to create project")
		return nil, err
	}

	return project, nil

}

func (s *projectService) GetProjectByProjectID(ctx context.Context, projectID uint) (*model.Project, error) {
	project, err := s.repo.GetProjectByProjectID(ctx, projectID)
	if err != nil {
		s.Log.WithError(err).Errorf("Failed to get project with projectID", projectID)
		return nil, err
	}
	return project, nil
}

func (s *projectService) GetProjectsByOwnerID(ctx context.Context, ownerID uint) ([]model.Project, error) {
	projects, err := s.repo.GetProjectsByOwnerID(ctx, ownerID)
	if err != nil {
		s.Log.WithError(err).Errorf("Failed to get projects for owner ID: %d", ownerID)
		return nil, err
	}
	return projects, nil
}
