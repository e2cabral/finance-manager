package routes

import (
	"finance-manager/src/presentation/controllers"
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	controller := controllers.AuthController{}
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/sign-up", controller.SignUp).Methods("POST")
}
