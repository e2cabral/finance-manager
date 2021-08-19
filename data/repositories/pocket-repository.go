package repositories

import (
	"finance-manager/domain/models"
	database "finance-manager/infra/database/client"
	"gorm.io/gorm"
)

type PocketRepository struct {
	handler *gorm.DB
}

func NewPocketRepository() (*PocketRepository, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &PocketRepository{handler: db}, nil
}

func (repository *PocketRepository) GetPockets(pockets *[]models.Pocket) {
	repository.handler.Find(&pockets)
}

func (repository *PocketRepository) GetPocketById(id uint, pocket *models.Pocket) {
	repository.handler.Find(&pocket, "id = ?", id)
}

func (repository *PocketRepository) Save(pocket *models.Pocket) {
	repository.handler.Create(&pocket)
}

func (repository *PocketRepository) Update(id uint, pocket *models.Pocket) {
	repository.handler.Updates(map[string]string{
		"type":        pocket.Type,
		"name":        pocket.Name,
		"description": pocket.Description,
	}).Where("id = ?", id)
}

func (repository *PocketRepository) Delete(id uint) {
	repository.handler.Delete(models.Pocket{}, "id = ?", id)
}
