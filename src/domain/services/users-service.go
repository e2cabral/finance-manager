package services

import (
	"finance-manager/src/data/repositories"
	"finance-manager/src/domain/models"
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

func (u *UsersService) Update(id uint, user models.User) (*models.User, error) {
	repository, err := repositories.NewUsersRepository()
	if err != nil {
		return nil, err
	}

	repository.Update(id, &user)
	return &user, nil
}

func (u *UsersService) Delete(id uint) error {
	repository, err := repositories.NewUsersRepository()
	if err != nil {
		return err
	}

	repository.Delete(id)
	return nil
}
