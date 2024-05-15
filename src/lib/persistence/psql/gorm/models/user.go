package models

import (
	"gym-management/src/lib/persistence/psql/gorm/customer_types"
	"time"
)

type User struct {
	Id         string               `gorm:"primaryKey"`
	Usernames  []Username           `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Password   string               `gorm:"not null"`
	Role       string               `gorm:"not null"`
	Profile    customer_types.JSONB `gorm:"type:jsonb;not null"`
	Restricted bool                 `gorm:"not null"`
	LastLogin  *time.Time           `gorm:"null"`
	//CreatedBy     string         `gorm:"not null"`
	//UpdatedBy     string         `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index;null"`
	//DeletedBy     *string        `gorm:"null"`
	//DeletedByUser *User          `gorm:"foreignKey:DeletedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedAt time.Time `gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null; default:CURRENT_TIMESTAMP"`
}
