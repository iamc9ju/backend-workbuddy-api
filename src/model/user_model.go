package model

import (
	"time"
)

type User struct {
	UserID   int    `gorm:"primaryKey;not null" json:"user_id"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Username string `gorm:"coulmn:username;not null" json:"username"`
	// FullName          string    `gorm:"" json:"full_name"`
	Phone             string    `gorm:"column:phone_number" json:"phone_number"`
	Role              string    `gorm:"column:role;default:developer;not null" json:"role"`
	ProfilePictureUrl string    `gorm:"" json:"profile_picture_url"`
	GoogleID          string    `gorm:"column:google_id" json:"google_id"`
	Location          string    `gorm:"" json:"location"` //ค่อยคิดว่าจะทำอะไรต่อ
	Gender            string    `gorm:"default:male not null json:"gender"`
	Nationality       string    `gorm:"default:thai not null" json:"nationality"`
	Bio               string    `gorm:"" json:"bio"`
	CreatedAt         time.Time `gorm:"autoCreateTime:milli" json:"create_at"`
	UpdatedAt         time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"update_at"`
}

// func (user *User) BeforeCreate(_ *gorm.DB) error {
// 	user.ID = uuid.New() // Generate UUID before create
// 	return nil
// }
