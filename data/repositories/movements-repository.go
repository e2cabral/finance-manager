package repositories

import (
	"finance-manager/domain/models"
	database "finance-manager/infra/database/client"
	"gorm.io/gorm"
)

type MovementRepository struct {
	handler *gorm.DB
}

func NewMovementRepository() (*MovementRepository, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &MovementRepository{handler: db}, nil
}

func (repository *MovementRepository) GetMovements(movements *[]models.Movement) {
	repository.handler.Find(&movements)
}

func (repository *MovementRepository) GetById(id uint, movements *models.Movement) {
	repository.handler.Find(&movements, "id = ?", id)
}

func (repository *MovementRepository) Save(movement *models.Movement) {
	repository.handler.Create(&movement)
}

func (repository *MovementRepository) Update(id uint, movement *models.Movement) {
	repository.handler.Update("value", &movement).Where("id = ?", id)
}
