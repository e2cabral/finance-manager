package controllers

import (
	"encoding/json"
	"finance-manager/domain/models"
	"finance-manager/domain/services"
	h "finance-manager/infra/helpers/http"
	"net/http"
)

type UsersController struct{}

func (controller *UsersController) Save(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.InternalServerError(w, err.Error())
	}

	service := services.UsersService{}
	newUser, err := service.Save(user)
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, newUser)
}
