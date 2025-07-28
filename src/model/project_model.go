package model

import "time"

type Project struct {
	ProjectID   uint      `gorm:"column:project_id;primaryKey" json:"project_id"`
	OwnerID     uint      `gorm:"column:owner_id;not null" json:"owner_id"`
	Title       string    `gorm:"column:title;size:255;not null" json:"title"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	Budget      float64   `gorm:"column:budget;type:decimal(16,2)" json:"budget"`
	Currency    string    `gorm:"column:currency;size:3;default:'THB'" json:"currency"`
	Deadline    time.Time `gorm:"column:deadline" json:"deadline"`
	Status      string    `gorm:"column:status;size:20;default:'draft'" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type ProjectCreateInput struct {
	OwnerID     uint      `json:"owner_id" validate:"required,numeric"`
	Title       string    `json:"title" validate:"required,max=255"`
	Description string    `json:"description" validate:"max=2000"`
	Budget      float64   `json:"budget" validate:"gt=0"`
	Currency    string    `json:"currency" validate:"required,len=3,uppercase"`
	Deadline    time.Time `json:"deadline" `
	Status      string    `json:"status" validate:"omitempty,oneof=draft open in_progress completed cancelled"`
}

type ProjectFilter struct {
}
