package models

import (
	"time"
)

type Gym struct {
	Id            string     `gorm:"primaryKey"`
	Name          string     `gorm:"not null"`
	Address       string     `gorm:"not null"`
	Enabled       bool       `gorm:"not null"`
	DisabledFor   *string    `gorm:"null"`
	OwnerId       string     `gorm:"not null"`
	Owner         GymOwner   `gorm:"foreignKey:OwnerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedBy     string     `gorm:"not null"`
	CreatedByUser User       `gorm:"foreignKey:CreatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UpdatedBy     string     `gorm:"not null"`
	UpdatedByUser User       `gorm:"foreignKey:UpdatedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DeletedBy     *string    `gorm:"null"`
	DeletedByUser *User      `gorm:"foreignKey:DeletedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedAt     time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:"null;index"`
}
