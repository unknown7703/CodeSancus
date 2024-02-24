package models

import "time"

type Datasets struct {
	name                string
	programmingLanguage string
	fileCount           int
	createdAt           time.Time `gorm:"not null"`
	updatedAt           time.Time `gorm:"not null"`
	token               string    `gorm:"not null;index:indexDatasetsOnToken"`
}
