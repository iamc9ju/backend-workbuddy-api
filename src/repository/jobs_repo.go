package repository

import (
	"app/src/model"
	"errors"
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type JobRepository interface {
	GetJobList(ctx context.Context) ([]model.Job, error)
	CreateJob(ctx context.Context, job *model.Job) error
	GetJobByJobID(ctx context.Context, jobID uint) (*model.Job, error)
	GetJobBySlug(ctx context.Context, slug string) (*model.Job, error)
	GetJobByOwnerID(ctx context.Context, ownerID uint) ([]model.Job, error)
}

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{db: db}
}

func (r *jobRepository) GetJobList(ctx context.Context) ([]model.Job, error) {
	var jobs []model.JobWithColor

	err := r.db.WithContext(ctx).
		Table("jobs").
		Select("jobs.*, colors.color_name, colors.color_code,pts.language_id,pts.framework_id").
		Joins("LEFT JOIN colors ON colors.color_id = jobs.background_color_id").
		Joins("LEFT JOIN project_tech_stack as pts on pts.project_id =jobs.job_id").
		Scan(&jobs).Error

	if err != nil {
		return nil, err
	}

	var result []model.Job
	for _, p := range jobs {
		job := p.Job
		job.LanguageId = p.LanguageId
		job.FrameWorkId = p.FrameWorkId
		job.Color = model.Color{
			ColorID:   p.BackgroundColorId,
			ColorName: p.ColorName,
			ColorCode: p.ColorCode,
		}
		result = append(result, job)
	}
	return result, nil
}

func (r *jobRepository) CreateJob(ctx context.Context, job *model.Job) error {
	return r.db.WithContext(ctx).Create(job).Error
}

func (r *jobRepository) GetJobByJobID(ctx context.Context, jobID uint) (*model.Job, error) {
	var project model.Job
	err := r.db.WithContext(ctx).
		Where("job_id = ?", jobID).
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

func (r *jobRepository) GetJobBySlug(ctx context.Context, slug string) (*model.Job, error) {
	var project model.Job
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

func (r *jobRepository) GetJobByOwnerID(ctx context.Context, ownerID uint) ([]model.Job, error) {
	var projects []model.Job
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
