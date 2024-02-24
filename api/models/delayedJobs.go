package models

import "time"

type DelayedJobs struct {
	priority  int    `gorm:"not null;default:0;index:delayedJobsPriority"`
	attempts  int    `gorm:"not null;default:0"`
	handler   string `gorm:"not null"`
	lastError string
	runAt     time.Time `gorm:"index:delayedJobsPriority"`
	lockedAt  time.Time
	failedAt  time.Time
	lockedBy  string
	queue     string
	createdAt time.Time
	updatedAt time.Time
}
