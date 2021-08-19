package routes

import (
	"finance-manager/presentation/controllers"
	"github.com/gorilla/mux"
)

func UsersRoutes(r *mux.Router) {
	controller := controllers.NewUsersController()
	r.HandleFunc("/users", controller.Save).Methods("POST")
	r.HandleFunc("/users/{id}", controller.Update).Methods("PATCH")
	r.HandleFunc("/users/{id}", controller.Delete).Methods("DELETE")
}
