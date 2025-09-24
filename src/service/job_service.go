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
	GetJobList(ctx context.Context) ([]model.Job, error)
	CreateJob(ctx context.Context, body model.JobCreate, ownderID uint) (*model.Job, error)
	GetJobByJobID(ctx context.Context, jobID uint) (*model.Job, error)
	GetJobBySlug(ctx context.Context, slug string) (*model.Job, error)
	GetJobByOwnerID(ctx context.Context, ownerID uint) ([]model.Job, error)
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

func (s *jobService) GetJobList(ctx context.Context) ([]model.Job, error) {
	s.Log.Info("Retrieving project list")
	projects, err := s.repo.GetJobList(ctx)
	if err != nil {
		s.Log.WithError(err).Error("Failed to retrieve project list")
		return nil, err
	}
	s.Log.Infof("Successfully retrieved %d projects", len(projects))
	return projects, nil
}

func (s *jobService) CreateJob(ctx context.Context, body model.JobCreate, ownerID uint) (*model.Job, error) {
	if err := s.Validate.Struct(body); err != nil {
		s.Log.WithError(err).Error("Validation failed")
		return nil, err
	}

	project := &model.Job{
		OwnerID:     body.OwnerID,
		Title:       body.Title,
		Description: body.Description,
		BudgetMin:   body.BudgetMin,
		BudgetMax:   body.BudgetMax,
		Currency:    body.Currency,
		Deadline:    body.Deadline,
		Status:      enum.DRAFT_PROJECT, // หรือ "open" ตาม business logic
	}

	if err := s.repo.CreateJob(ctx, project); err != nil {
		s.Log.WithError(err).Error("Failed to create project")
		return nil, err
	}

	return project, nil

}

func (s *jobService) GetJobByJobID(ctx context.Context, jobID uint) (*model.Job, error) {
	project, err := s.repo.GetJobByJobID(ctx, jobID)
	if err != nil {
		s.Log.WithError(err).Errorf("Failed to get job with project ID: %d", jobID)
		return nil, err
	}
	return project, nil
}

func (s *jobService) GetJobBySlug(ctx context.Context, slug string) (*model.Job, error) {
	project, err := s.repo.GetJobBySlug(ctx, slug)
	if err != nil {
		s.Log.WithError(err).Errorf("Failed to get job with job slug: %d", slug)
		return nil, err
	}
	return project, nil
}

func (s *jobService) GetJobByOwnerID(ctx context.Context, ownerID uint) ([]model.Job, error) {
	projects, err := s.repo.GetJobByOwnerID(ctx, ownerID)
	if err != nil {
		s.Log.WithError(err).Errorf("Failed to get jobs for owner ID: %d", ownerID)
		return nil, err
	}
	return projects, nil
}
