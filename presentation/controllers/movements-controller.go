package controllers

import (
	"encoding/json"
	"finance-manager/domain/models"
	"finance-manager/domain/services"
	h "finance-manager/infra/helpers/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MovementController struct{}

func NewMovementsController() *MovementController {
	return &MovementController{}
}

func (m *MovementController) GetMovements(w http.ResponseWriter, r *http.Request) {
	service := services.MovementService{}

	movements, err := service.GetMovements()
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, movements)
}

func (m *MovementController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := services.MovementService{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	movement, err := service.GetById(uint(id))
	h.Ok(w, movement)
}

func (m *MovementController) Save(w http.ResponseWriter, r *http.Request) {
	var movement models.Movement

	err := json.NewDecoder(r.Body).Decode(&movement)

	service := services.MovementService{}
	newMovement, err := service.Save(movement)
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, newMovement)
}

func (m *MovementController) Update(w http.ResponseWriter, r *http.Request) {
	var movement models.Movement

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	err = json.NewDecoder(r.Body).Decode(&movement)
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	service := services.MovementService{}
	updated, err := service.Update(uint(id), movement)

	h.Ok(w, updated)
}
