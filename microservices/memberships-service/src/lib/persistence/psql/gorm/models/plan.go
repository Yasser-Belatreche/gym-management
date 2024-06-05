package models

import (
	"time"
)

type Plan struct {
	Id              string     `gorm:"primaryKey"`
	Name            string     `gorm:"not null"`
	Featured        bool       `gorm:"not null"`
	SessionsPerWeek int        `gorm:"not null"`
	WithCoach       bool       `gorm:"not null"`
	MonthlyPrice    float64    `gorm:"not null"`
	GymId           string     `gorm:"not null"`
	CreatedBy       string     `gorm:"not null"`
	UpdatedBy       string     `gorm:"not null"`
	DeletedBy       *string    `gorm:"null"`
	DeletedAt       *time.Time `gorm:"null;index"`
	CreatedAt       time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
}
