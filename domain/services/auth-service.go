package services

import (
	"finance-manager/data/repositories"
	"finance-manager/domain/models"
)

type AuthService struct{}

func (a *AuthService) Login(username string, password string) (*models.User, error) {
	var user models.User
	repository, err := repositories.NewUsersRepository()
	if err != nil {
		return nil, err
	}

	repository.GetUserByUsername(username, password, &user)
	return &user, nil
}
