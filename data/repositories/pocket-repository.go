package repositories

import (
	"finance-manager/domain/models"
	database "finance-manager/infra/database/client"
	"gorm.io/gorm"
)

type PocketRepository struct {
	Handler *gorm.DB
}

func NewPocketRepository() (*PocketRepository, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &PocketRepository{Handler: db}, nil
}

func (repository *PocketRepository) GetPockets(pockets *[]models.Pocket) {
	repository.Handler.Find(&pockets)
}

func (repository *PocketRepository) GetPocketById(id uint, pocket *models.Pocket) {
	repository.Handler.Find(&pocket, "id = ?", id)
}

func (repository *PocketRepository) Save(pocket *models.Pocket) {
	repository.Handler.Create(&pocket)
}

func (repository *PocketRepository) Update(id uint, pocket *models.Pocket) {
	repository.Handler.Updates(map[string]string{
		"type":        pocket.Type,
		"name":        pocket.Name,
		"description": pocket.Description,
	}).Where("id = ?", id)
}

func (repository *PocketRepository) Delete(id uint) {
	repository.Handler.Delete(models.Pocket{}, "id = ?", id)
}
