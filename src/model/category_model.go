package model

import "time"

type Category struct {
	CategoryID  uint      `gorm:"column:category_id;primaryKey" json:"category_id"`
	Title       string    `gorm:"column:title;size:255;not null" json:"title"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	Status      string    `gorm:"column:status;size:20;default:'ACTIVE'" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type CategoryCreate struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description" validate:"max=2000"`
	Status      string `json:"status" validate:"omitempty,oneof=ACTIVE open in_progress completed cancelled"`
}

type CategoryFilter struct {
}
