package models

import "time"

type Reports struct {
	datasetId  int `gorm:"not null;index:indexReportsOnDatasetId"`
	status     int
	err        string
	stdout     string
	stderr     string
	exitStatus int
	memory     int
	runTime    float32
	createdAt  time.Time `gorm:"not null"`
	updatedAt  time.Time `gorm:"not null"`
	token      string    `gorm:"not null;index:indexReportsOnToken"`
}
