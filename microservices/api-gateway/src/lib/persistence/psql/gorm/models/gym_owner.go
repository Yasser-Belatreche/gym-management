package models

import (
	"time"
)

type GymOwner struct {
	Id            string     `gorm:"primaryKey"`
	Name          string     `gorm:"not null"`
	PhoneNumber   string     `gorm:"not null"`
	Email         string     `gorm:"not null"`
	Restricted    bool       `gorm:"not null"`
	CreatedBy     string     `gorm:"not null"`
	Gyms          []Gym      `gorm:"foreignKey:OwnerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedByUser User       `gorm:"foreignKey:CreatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UpdatedBy     string     `gorm:"not null"`
	UpdatedByUser User       `gorm:"foreignKey:UpdatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DeletedBy     *string    `gorm:"null"`
	DeletedByUser *User      `gorm:"foreignKey:DeletedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DeletedAt     *time.Time `gorm:"null;index"`
	CreatedAt     time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
}
