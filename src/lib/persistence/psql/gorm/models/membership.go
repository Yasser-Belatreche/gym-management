package models

import (
	"time"
)

type Membership struct {
	Id              string     `gorm:"primaryKey"`
	Code            string     `gorm:"unique;not null"`
	StartDate       time.Time  `gorm:"not null"`
	EndDate         *time.Time `gorm:"null"`
	Enabled         bool       `gorm:"not null"`
	DisabledFor     *string    `gorm:"null"`
	SessionsPerWeek int        `gorm:"not null"`
	WithCoach       bool       `gorm:"not null"`
	MonthlyPrice    float64    `gorm:"not null"`
	PlanId          string     `gorm:"not null"`
	Plan            Plan       `gorm:"foreignKey:PlanId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CustomerId      string     `gorm:"not null"`
	Customer        Customer   `gorm:"foreignKey:CustomerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	RenewedAt       *time.Time `gorm:"null"`
	CreatedBy       string     `gorm:"not null"`
	CreatedByUser   User       `gorm:"foreignKey:CreatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UpdatedBy       string     `gorm:"not null"`
	UpdatedByUser   User       `gorm:"foreignKey:UpdatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedAt       time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
}
