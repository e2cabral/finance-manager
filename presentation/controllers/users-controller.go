package controllers

import (
	"encoding/json"
	"finance-manager/domain/models"
	"finance-manager/domain/services"
	h "finance-manager/infra/helpers/http"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UsersController struct{}

func NewUsersController() *UsersController {
	return &UsersController{}
}

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

func (controller *UsersController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.InternalServerError(w, err.Error())
	}

	service := services.UsersService{}
	updated, err := service.Update(uint(id), user)
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, updated)
}

func (controller *UsersController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	service := services.UsersService{}
	err = service.Delete(uint(id))
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, fmt.Sprintf("The user with id %d, was deleted.", id))
}
