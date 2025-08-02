package repository

import (
	"app/src/model"
	"errors"
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetProjectList(ctx context.Context) ([]model.Project, error)
	CreateProject(ctx context.Context, project *model.Project) error
	GetProjectByProjectID(ctx context.Context, projectID uint) (*model.Project, error)
	GetProjectBySlug(ctx context.Context, slug string) (*model.Project, error)
	GetProjectsByOwnerID(ctx context.Context, ownerID uint) ([]model.Project, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) GetProjectList(ctx context.Context) ([]model.Project, error) {
	var projects []model.Project

	err := r.db.WithContext(ctx).Find(&projects).Error

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) CreateProject(ctx context.Context, project *model.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

func (r *projectRepository) GetProjectByProjectID(ctx context.Context, projectID uint) (*model.Project, error) {
	var project model.Project
	err := r.db.WithContext(ctx).
		Where("project_id = ?", projectID).
		First(&project).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("project not found or you don't have permission")
		}
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) GetProjectBySlug(ctx context.Context, slug string) (*model.Project, error) {
	var project model.Project
	err := r.db.WithContext(ctx).
		Where("slug = ?", slug).
		First(&project).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("project not found or you don't have permission")
		}
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) GetProjectsByOwnerID(ctx context.Context, ownerID uint) ([]model.Project, error) {
	var projects []model.Project
	err := r.db.WithContext(ctx).
		Where("owner_id = ?", ownerID).
		Find(&projects).
		Error

	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("no projects found for this owner")
	}

	return projects, nil
}
