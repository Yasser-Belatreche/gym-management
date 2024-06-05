package models

type Username struct {
	UserId   string `gorm:"primaryKey;"`
	Username string `gorm:"unique;primaryKey"`
}
