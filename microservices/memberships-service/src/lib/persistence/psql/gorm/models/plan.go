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
	Gym             Gym        `gorm:"foreignKey:GymId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedBy       string     `gorm:"not null"`
	CreatedByUser   User       `gorm:"foreignKey:CreatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UpdatedBy       string     `gorm:"not null"`
	UpdatedByUser   User       `gorm:"foreignKey:UpdatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DeletedBy       *string    `gorm:"null"`
	DeletedByUser   *User      `gorm:"foreignKey:DeletedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DeletedAt       *time.Time `gorm:"null;index"`
	CreatedAt       time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
}
