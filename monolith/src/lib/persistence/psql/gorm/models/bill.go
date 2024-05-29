package models

import (
	"time"
)

type Bill struct {
	Id           string     `gorm:"primaryKey"`
	Amount       float64    `gorm:"not null"`
	Paid         bool       `gorm:"not null"`
	PaidAt       *time.Time `gorm:"null"`
	DueTo        time.Time  `gorm:"not null"`
	MembershipId string     `gorm:"not null"`
	Membership   Membership `gorm:"foreignKey:MembershipId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedAt    time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time  `gorm:"not null; default:CURRENT_TIMESTAMP"`
}
