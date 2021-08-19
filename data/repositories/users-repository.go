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

func (repository *UsersRepository) Update(id uint, user *models.User) {
	repository.handler.Update("username, password", &user).Where("id = ?", id)
}

func (repository *UsersRepository) Delete(id uint) {
	repository.handler.Delete(models.User{}, "id = ?", id)
}
