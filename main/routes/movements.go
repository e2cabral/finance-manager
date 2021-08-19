package routes

import (
	h "finance-manager/infra/helpers/auth"
	"finance-manager/presentation/controllers"
	"github.com/gorilla/mux"
)

func MovementsRoutes(r *mux.Router) {
	controller := controllers.NewMovementsController()
	r.HandleFunc("/movements/{pocket_id}", h.IsAuthenticated(controller.GetMovements)).Methods("GET")
	r.HandleFunc("/movements/{pocket_id}/{id}", h.IsAuthenticated(controller.GetById)).Methods("GET")
	r.HandleFunc("/movements/{pocket_id}", h.IsAuthenticated(controller.Save)).Methods("POST")
	r.HandleFunc("/movements/{pocket_id}/{id}", h.IsAuthenticated(controller.Update)).Methods("PATCH")
	r.HandleFunc("/movements/{pocket_id}/{id}", h.IsAuthenticated(controller.Update)).Methods("DELETE")
}
