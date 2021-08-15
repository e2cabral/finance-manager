package services

import (
	"finance-manager/data/repositories"
	"finance-manager/domain/models"
)

type PocketService struct{}

func (service *PocketService) GetPockets() (*[]models.Pocket, error) {
	var pockets []models.Pocket

	repository, err := repositories.NewPocketRepository()
	if err != nil {
		return nil, err
	}

	repository.GetPockets(&pockets)
	return &pockets, nil
}

func (service *PocketService) GetPocketById(id uint) (*models.Pocket, error) {
	var pocket models.Pocket

	repository, err := repositories.NewPocketRepository()
	if err != nil {
		return nil, err
	}

	repository.GetPocketById(id, &pocket)
	return &pocket, nil
}

func (service *PocketService) Save(pocket models.Pocket) (*models.Pocket, error) {
	repository, err := repositories.NewPocketRepository()
	if err != nil {
		return nil, err
	}

	repository.Save(&pocket)
	return &pocket, nil
}

func (service *PocketService) Update(id uint, pocket models.Pocket) (*models.Pocket, error) {
	repository, err := repositories.NewPocketRepository()
	if err != nil {
		return nil, err
	}

	repository.Update(id, &pocket)
	return &pocket, nil
}
