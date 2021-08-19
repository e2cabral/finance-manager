package services

import (
	"finance-manager/data/repositories"
	"finance-manager/domain/models"
)

type MovementService struct{}

func (m *MovementService) GetMovements(pocketId uint) ([]models.Movement, error) {
	var movements []models.Movement

	repository, err := repositories.NewMovementRepository()
	if err != nil {
		return nil, err
	}

	repository.GetMovements(pocketId, &movements)
	return movements, nil
}

func (m *MovementService) GetById(id uint, pocketId uint) (*models.Movement, error) {
	var movement models.Movement

	repository, err := repositories.NewMovementRepository()
	if err != nil {
		return nil, err
	}

	repository.GetById(id, pocketId, &movement)
	return &movement, nil
}

func (m *MovementService) Save(pocketId uint, movement models.Movement) (*models.Movement, error) {
	repository, err := repositories.NewMovementRepository()
	if err != nil {
		return nil, err
	}

	movement.PocketID = pocketId
	repository.Save(&movement)
	return &movement, nil
}

func (m *MovementService) Update(id uint, pocketId uint, movement models.Movement) (*models.Movement, error) {
	repository, err := repositories.NewMovementRepository()
	if err != nil {
		return nil, err
	}

	repository.Update(id, pocketId, &movement)
	return &movement, nil
}

func (m *MovementService) Delete(id uint) error {
	repository, err := repositories.NewMovementRepository()
	if err != nil {
		return err
	}

	repository.Delete(id)
	return nil
}
