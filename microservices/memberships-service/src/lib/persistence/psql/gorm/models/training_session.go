package models

import (
	"time"
)

type TrainingSession struct {
	Id           string     `gorm:"primaryKey"`
	MembershipId string     `gorm:"not null"`
	Membership   Membership `gorm:"foreignKey:MembershipId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	StartedAt    time.Time  `gorm:"not null"`
	EndedAt      *time.Time `gorm:"null"`
}
