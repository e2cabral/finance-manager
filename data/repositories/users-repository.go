package repositories

import (
	"finance-manager/domain/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	handler *gorm.DB
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (repository *UsersRepository) Create(user *models.User) {
	repository.handler.Create(&user)
}
