package model

import (
	"time"
)

type Job struct {
	JobID             uint      `gorm:"column:job_id;primaryKey" json:"job_id"`
	OwnerID           uint      `gorm:"column:owner_id;not null" json:"owner_id"`
	CategoryId        uint      `gorm:"column:category_id;not null" json:"category_id"`
	Title             string    `gorm:"column:title;size:255;not null" json:"title"`
	Slug              string    `gorm:"column:slug;size:255;not null" json:"slug"`
	Description       string    `gorm:"column:description;type:text" json:"description"`
	BudgetMin         float64   `gorm:"column:budget_min;type:decimal(16,2)" json:"budget_min"`
	BudgetMax         float64   `gorm:"column:budget_max;type:decimal(16,2)" json:"budget_max"`
	Currency          string    `gorm:"column:currency;size:3;default:'THB'" json:"currency"`
	Deadline          time.Time `gorm:"column:deadline" json:"deadline"`
	Status            string    `gorm:"column:status;size:20;default:'DRAFT'" json:"status"`
	BackgroundColorId int       `gorm:"column:background_color_id" json:"background_color_id"`
	LanguageId        int       `gorm:"column:language_id" json:"language_id"`
	FrameWorkId       int       `gorm:"column:framework_id" json:"frameword_id"`
	Color             Color     `gorm:"foreignKey:BackgroundColorId;references:ColorID" json:"color"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
type JobWithColor struct {
	Job
	ColorName string `gorm:"column:color_name" json:"color_name"`
	ColorCode string `gorm:"column:color_code" json:"color_code"`
}
type JobCreate struct {
	OwnerID     uint      `json:"owner_id" validate:"required,numeric"`
	Title       string    `json:"title" validate:"required,max=255"`
	Description string    `json:"description" validate:"max=2000"`
	BudgetMin   float64   `gorm:"column:budget_min;type:decimal(16,2)" json:"budget_min"`
	BudgetMax   float64   `gorm:"column:budget_max;type:decimal(16,2)" json:"budget_max"`
	Currency    string    `json:"currency" validate:"required,len=3,uppercase"`
	Deadline    time.Time `json:"deadline" `
	Status      string    `json:"status" validate:"omitempty,oneof=DRAFT open in_progress completed cancelled"`
}

type Color struct {
	ColorID   int    `gorm:"column:color_id;primaryKey" json:"color_id"`
	ColorName string `gorm:"column:color_name" json:"color_name"`
	ColorCode string `gorm:"column:color_code" json:"color_code"`
}

type JobFilter struct {
}
