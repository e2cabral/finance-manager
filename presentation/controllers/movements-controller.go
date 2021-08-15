package controllers

import (
	"encoding/json"
	"finance-manager/domain/models"
	"finance-manager/domain/services"
	h "finance-manager/infra/helpers/http"
	"net/http"
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