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

type MovementController struct{}

func NewMovementsController() *MovementController {
	return &MovementController{}
}

func (m *MovementController) GetMovements(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pocketId, err := strconv.Atoi(vars["pocket_id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	service := services.MovementService{}

	movements, err := service.GetMovements(uint(pocketId))
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

	pocketId, err := strconv.Atoi(vars["pocket_id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	movement, err := service.GetById(uint(id), uint(pocketId))
	h.Ok(w, movement)
}

func (m *MovementController) Save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movement models.Movement
	service := services.MovementService{}

	pocketId, err := strconv.Atoi(vars["pocket_id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	if err := json.NewDecoder(r.Body).Decode(&movement); err != nil {
		h.InternalServerError(w, err.Error())
	}

	newMovement, err := service.Save(uint(pocketId), movement)
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, newMovement)
}

func (m *MovementController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movement models.Movement

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	pocketId, err := strconv.Atoi(vars["pocket_id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	err = json.NewDecoder(r.Body).Decode(&movement)
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	service := services.MovementService{}
	updated, err := service.Update(uint(id), uint(pocketId), movement)

	h.Ok(w, updated)
}

func (m *MovementController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := services.MovementService{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	pocketId, err := strconv.Atoi(vars["pocket_id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	err = service.Delete(uint(id), uint(pocketId))
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, fmt.Sprintf("Movement with id %d, was deleted.", id))
}
