package database

import (
	"finance-manager/domain/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect() (*gorm.DB, error) {
	password := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USERNAME")
	//defaultDb := os.Getenv("DEFAULT_DB")
	host := os.Getenv("DB_HOST")
	table := os.Getenv("DB_TABLE")
	port := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, table, port,
	)

	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return database, nil
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Movement{}, &models.Pocket{}); err != nil {
		return err
	}
	return nil
}
