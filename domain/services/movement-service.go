package services

import (
	"finance-manager/data/repositories"
	"finance-manager/domain/models"
)

type MovementService struct{}

func (m *MovementService) GetMovements() ([]models.Movement, error) {
	var movements []models.Movement

	repository, err := repositories.NewMovementRepository()
	if err != nil {
		return nil, err
	}

	repository.GetMovements(&movements)
	return movements, nil
}

func (m *MovementService) GetById(id uint) (models.Movement, error) {
	var movement models.Movement

	repository, err := repositories.NewMovementRepository()
	if err != nil {
		return models.Movement{}, err
	}

	repository.GetById(id, &movement)
	return movement, nil
}

func (m *MovementService) Save(movement models.Movement) (models.Movement, error) {
	var newMovement models.Movement

	repository, err := repositories.NewMovementRepository()
	if err != nil {
		return models.Movement{}, err
	}

	repository.Save(&movement)
	return newMovement, nil
}
