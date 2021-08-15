package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	connectionString := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return database, nil
}
