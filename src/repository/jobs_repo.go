package repository

import (
	"app/src/model"
	"errors"
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type JobRepository interface {
	GetProjectList(ctx context.Context) ([]model.Project, error)
	CreateProject(ctx context.Context, project *model.Project) error
	GetProjectByProjectID(ctx context.Context, projectID uint) (*model.Project, error)
	GetProjectBySlug(ctx context.Context, slug string) (*model.Project, error)
	GetProjectsByOwnerID(ctx context.Context, ownerID uint) ([]model.Project, error)
}

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{db: db}
}

func (r *jobRepository) GetProjectList(ctx context.Context) ([]model.Project, error) {
	var projects []model.ProjectWithColor

	err := r.db.WithContext(ctx).
		Table("projects").
		Select("projects.*, colors.color_name, colors.color_code,pts.language_id,pts.framework_id").
		Joins("LEFT JOIN colors ON colors.color_id = projects.background_color_id").
		Joins("LEFT JOIN project_tech_stack as pts on pts.project_id =projects.project_id").
		Scan(&projects).Error

	if err != nil {
		return nil, err
	}

	var result []model.Project
	for _, p := range projects {
		project := p.Project
		project.LanguageId = p.LanguageId
		project.FrameWorkId = p.FrameWorkId
		project.Color = model.Color{
			ColorID:   p.BackgroundColorId,
			ColorName: p.ColorName,
			ColorCode: p.ColorCode,
		}
		result = append(result, project)
	}
	return result, nil
}

func (r *jobRepository) CreateProject(ctx context.Context, project *model.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

func (r *jobRepository) GetProjectByProjectID(ctx context.Context, projectID uint) (*model.Project, error) {
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

func (r *jobRepository) GetProjectBySlug(ctx context.Context, slug string) (*model.Project, error) {
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

func (r *jobRepository) GetProjectsByOwnerID(ctx context.Context, ownerID uint) ([]model.Project, error) {
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
