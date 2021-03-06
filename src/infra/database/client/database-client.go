package database

import (
	models2 "finance-manager/src/domain/models"
	"finance-manager/src/infra/environment"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	env := environment.Env{}
	password, err := env.GetVariable("DB_PASSWORD")
	username, err := env.GetVariable("DB_USERNAME")
	host, err := env.GetVariable("DB_HOST")
	table, err := env.GetVariable("DB_TABLE")
	port, err := env.GetVariable("DB_PORT")

	if err != nil {
		return nil, err
	}

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
	if err := db.AutoMigrate(&models2.Movement{}, &models2.Pocket{}, &models2.User{}); err != nil {
		return err
	}
	return nil
}
