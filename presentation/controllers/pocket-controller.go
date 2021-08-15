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

type PocketController struct{}

func NewPocketController() *PocketController {
	return &PocketController{}
}

func (controller *PocketController) GetPockets(w http.ResponseWriter, r *http.Request) {
	service := services.PocketService{}

	pockets, err := service.GetPockets()
	if err != nil {
		h.InternalServerError(w, err.Error())
	}
	h.Ok(w, &pockets)
}

func (controller *PocketController) GetPocketById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	service := services.PocketService{}

	pocket, err := service.GetPocketById(uint(id))
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, &pocket)
}

func (controller *PocketController) Save(w http.ResponseWriter, r *http.Request) {
	var pocket models.Pocket
	if err := json.NewDecoder(r.Body).Decode(&pocket); err != nil {
		h.InternalServerError(w, err.Error())
	}

	service := services.PocketService{}
	newPocket, err := service.Save(pocket)
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, newPocket)
}

func (controller *PocketController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var pocket models.Pocket
	service := services.PocketService{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	if err := json.NewDecoder(r.Body).Decode(&pocket); err != nil {
		h.InternalServerError(w, err.Error())
	}

	updated, err := service.Update(uint(id), pocket)
	if err != nil {
		h.InternalServerError(w, err.Error())
	}

	h.Ok(w, &updated)
}
