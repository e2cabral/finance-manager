package routes

import (
	"finance-manager/presentation/controllers"
	"github.com/gorilla/mux"
)

func MovementsRoutes(r *mux.Router) {
	controller := controllers.NewMovementsController()
	r.HandleFunc("/movements/{pocket_id}", controller.GetMovements).Methods("GET")
	r.HandleFunc("/movements/{pocket_id}/{id}", controller.GetById).Methods("GET")
	r.HandleFunc("/movements/{pocket_id}", controller.Save).Methods("POST")
	r.HandleFunc("/movements/{pocket_id}/{id}", controller.Update).Methods("PATCH")
	r.HandleFunc("/movements/{pocket_id}/{id}", controller.Update).Methods("DELETE")
}
