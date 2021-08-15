package routes

import (
	"finance-manager/presentation/controllers"
	"github.com/gorilla/mux"
)

func MovementsRoutes(r *mux.Router) {
	controller := controllers.NewMovementsController()
	r.HandleFunc("/movements", controller.GetMovements).Methods("GET")
	r.HandleFunc("/movements/{id}", controller.GetById).Methods("GET")
	r.HandleFunc("/movements", controller.Save).Methods("POST")
	r.HandleFunc("/movements/{id}", controller.Update).Methods("PATCH")
}
