package controllers

import (
	"finance-manager/domain/services"
	h "finance-manager/infra/helpers/http"
	"net/http"
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
