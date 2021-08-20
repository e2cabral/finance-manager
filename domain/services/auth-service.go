package services

import (
	"finance-manager/data/repositories"
	"finance-manager/domain/models"
	"finance-manager/infra/helpers/encryption"
)

type AuthService struct{}

func (a *AuthService) Login(username string, password string) (*models.User, error) {
	var user models.User
	repository, err := repositories.NewUsersRepository()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := encryption.Encrypt(password)
	if err != nil {
		return nil, err
	}

	repository.GetUserByUsernameAndPassword(username, hashedPassword, &user)
	return &user, nil
}

func (a *AuthService) SignUp(user models.User) (*models.User, error) {
	repository, err := repositories.NewUsersRepository()
	if err != nil {
		return nil, err
	}

	user.Password, err = encryption.Encrypt(user.Password)
	if err != nil {
		return nil, err
	}

	repository.Save(&user)
	return &user, nil
}

func (a *AuthService) IsUsernameUsed(username string) (bool, error) {
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