package controllers

import (
	"finance-manager/domain/models"
	"finance-manager/domain/services"
	"finance-manager/infra/auth"
	h "finance-manager/infra/helpers/http"
	helpers "finance-manager/infra/helpers/json"
	"net/http"
)

type AuthController struct{}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := helpers.FromJSON(r.Body, &user); err != nil {
		h.InternalServerError(w, err.Error())
		return
	}

	service := services.AuthService{}

	u, err := service.Login(user.Username, user.Password)
	if err != nil {
		h.InternalServerError(w, err.Error())
		return
	}

	if u.ID == 0 {
		h.NotFound(w, "Credentials not valid.")
		return
	}

	token, err := auth.GenerateJWTToken(u.Username)
	if err != nil {
		h.InternalServerError(w, err.Error())
		return
	}

	authentication := models.Auth{
		Username: u.Username,
		Token:    token,
	}

	h.Ok(w, authentication)
}

func (a *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := helpers.FromJSON(r.Body, &user); err != nil {
		h.InternalServerError(w, err.Error())
		return
	}

	service := services.UsersService{}

	userExists, err := service.IsUsernameUsed(user.Username)
	if err != nil {
		h.InternalServerError(w, err.Error())
		return
	}

	if userExists {
		h.Ok(w, "Username is already used.")
		return
	}

	u, err := service.Save(user)
	if err != nil {
		h.InternalServerError(w, err.Error())
		return
	}

	h.Ok(w, u)
}
