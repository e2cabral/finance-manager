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

func (u *UsersService) IsUsernameUsed(username string) (bool, error) {
	var user models.User
	repository, err := repositories.NewUsersRepository()
	if err != nil {
		return false, err
	}

	repository.GetUserByUsername(username, &user)

	if &user != nil {
		return true, nil
	}
	return false, nil
}
