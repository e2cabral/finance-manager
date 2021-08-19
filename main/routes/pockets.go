package routes

import (
	"finance-manager/presentation/controllers"
	"github.com/gorilla/mux"
)

func PocketRoutes(r *mux.Router) {
	controller := controllers.NewPocketController()
	r.HandleFunc("/pockets", controller.GetPockets).Methods("GET")
	r.HandleFunc("/pockets/{id}", controller.GetPocketById).Methods("GET")
	r.HandleFunc("/pockets", controller.Save).Methods("POST")
	r.HandleFunc("/pockets/{id}", controller.Update).Methods("PATCH")
	r.HandleFunc("/pockets/{id}", controller.Delete).Methods("DELETE")
}
