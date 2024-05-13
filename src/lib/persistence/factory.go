package persistence

import (
	"gym-management/src/lib/persistence/psql/gorm"
	"os"
)

func InitializePersistence() {
	err := NewPersistence().Connect()
	if err != nil {
		panic(err)
	}
}

func NewPersistence() Persistence {
	return NewGormPersistence()
}

func NewGormPersistence() *gorm.GormPsqlPersistence {
	host, found := os.LookupEnv("DB_HOST")
	if !found {
		panic("DB_HOST env var is required")
	}

	user, found := os.LookupEnv("DB_USER")
	if !found {
		panic("DB_USER env var is required")
	}

	password, found := os.LookupEnv("DB_PASSWORD")
	if !found {
		panic("DB_PASSWORD env var is required")
	}

	database, found := os.LookupEnv("DB_NAME")
	if !found {
		panic("DB_NAME env var is required")
	}

	port, found := os.LookupEnv("DB_PORT")
	if !found {
		panic("DB_PORT env var is required")
	}

	return gorm.NewGormPsqlPersistence(host, user, password, database, port)
}
