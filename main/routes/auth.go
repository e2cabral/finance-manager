package routes

import (
	"finance-manager/presentation/controllers"
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	controller := controllers.AuthController{}
	r.HandleFunc("/login", controller.Login).Methods("POST")
}
