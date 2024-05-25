package models

import (
	"time"
)

type Customer struct {
	Id            string     `gorm:"primaryKey"`
	FirstName     string     `gorm:"not null"`
	LastName      string     `gorm:"not null"`
	PhoneNumber   string     `gorm:"not null"`
	Email         string     `gorm:"not null"`
	BirthYear     int        `gorm:"not null"`
	Gender        string     `gorm:"not null"`
	Restricted    bool       `gorm:"not null"`
	CreatedBy     string     `gorm:"not null"`
	CreatedByUser User       `gorm:"foreignKey:CreatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UpdatedBy     string     `gorm:"not null"`
	UpdatedByUser User       `gorm:"foreignKey:UpdatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DeletedBy     *string    `gorm:"null"`
	DeletedByUser *User      `gorm:"foreignKey:DeletedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedAt     time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:"null;index"`

	Memberships []Membership
}
