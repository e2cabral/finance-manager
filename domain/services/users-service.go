package services

import (
	"finance-manager/data/repositories"
	"finance-manager/domain/models"
)

type UsersService struct{}

func (u *UsersService) Save(user models.User) (*models.User, error) {
	repository, err := repositories.NewUsersRepository()
	if err != nil {
		return nil, err
	}

	repository.Save(&user)
	return &user, nil
}
