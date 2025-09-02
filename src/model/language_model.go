package model

type ProgrammingLanguage struct {
	LanguageID   int    `gorm:"column:language_id;primaryKey" json:"language_id"`
	LanguageName string `gorm:"column:language_name" json:"language_name"`
	// createAt
	// updateAt
}

type ProgrammingLanguageRequest struct {
	LanguageName string `gorm:"column:language_id" json:"language_name" binding:"required"`
}
