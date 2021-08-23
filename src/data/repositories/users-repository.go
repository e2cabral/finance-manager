package repositories

import (
	"finance-manager/src/domain/models"
	"finance-manager/src/infra/database/client"
	"gorm.io/gorm"
)

type UsersRepository struct {
	Handler *gorm.DB
}

func NewUsersRepository() (*UsersRepository, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &UsersRepository{
		Handler: db,
	}, nil
}

func (repository *UsersRepository) GetUserByUsernameAndPassword(username string, password string, user *models.User) {
	repository.Handler.Find(
		&user,
		"username = ? AND password = ?",
		username,
		password,
	)
}

func (repository *UsersRepository) GetUserByUsername(username string, user *models.User) {
	repository.Handler.Find(&user, "username = ?", username)
}

func (repository *UsersRepository) Save(user *models.User) {
	repository.Handler.Create(&user)
}

func (repository *UsersRepository) Update(id uint, user *models.User) {
	repository.Handler.Update("username, password", &user).Where("id = ?", id)
}

func (repository *UsersRepository) Delete(id uint) {
	repository.Handler.Delete(models.User{}, "id = ?", id)
}
