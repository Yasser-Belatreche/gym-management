package models

import (
	"time"
)

type Customer struct {
	Id          string     `gorm:"primaryKey"`
	FirstName   string     `gorm:"not null"`
	LastName    string     `gorm:"not null"`
	PhoneNumber string     `gorm:"not null"`
	Email       string     `gorm:"not null"`
	BirthYear   int        `gorm:"not null"`
	Gender      string     `gorm:"not null"`
	Restricted  bool       `gorm:"not null"`
	CreatedBy   string     `gorm:"not null"`
	UpdatedBy   string     `gorm:"not null"`
	DeletedBy   *string    `gorm:"null"`
	CreatedAt   time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `gorm:"null;index"`
}
