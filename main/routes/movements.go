package routes

import (
	"finance-manager/presentation/controllers"
	"github.com/gorilla/mux"
)

func MovementsRoutes(r *mux.Router) {
	controller := controllers.NewMovementsController()
	r.HandleFunc("/movements", controller.GetMovements).Methods("GET")
	r.HandleFunc("/movements", controller.Save).Methods("POST")
}
