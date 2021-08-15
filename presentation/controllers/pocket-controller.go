package controllers

import (
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
