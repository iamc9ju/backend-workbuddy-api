package model

import (
	"time"
)

type User struct {
	UserID            int       `gorm:"primaryKey;not null" json:"user_id"`
	Email             string    `gorm:"uniqueIndex;not null" json:"email"`
	Username          string    `gorm:"not null" json:"username"`
	FullName          string    `gorm:"" json:"full_name"`
	Phone             string    `gorm:"" json:"phone"`
	Role              string    `gorm:"default:developer;not null" json:"role"`
	ProfilePictureUrl string    `gorm:"" json:"profile_picture_url"`
	GoogleID          string    `gorm:"" json:"-"`
	Location          string    `gorm:"" json:"location"` //ค่อยคิดว่าจะทำอะไรต่อ
	Gender            string    `gorm:"default:male not null json:"gender"`
	Nationality       string    `gorm:"default:thai not null" json"nationality"`
	Bio               string    `gorm:"" json"bio"`
	CreatedAt         time.Time `gorm:"autoCreateTime:milli" json:"-"`
	UpdatedAt         time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"-"`
}

// func (user *User) BeforeCreate(_ *gorm.DB) error {
// 	user.ID = uuid.New() // Generate UUID before create
// 	return nil
// }
