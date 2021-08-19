package routes

import (
	h "finance-manager/infra/helpers/auth"
	"finance-manager/presentation/controllers"
	"github.com/gorilla/mux"
)

func UsersRoutes(r *mux.Router) {
	controller := controllers.NewUsersController()
	r.HandleFunc("/users", controller.Save).Methods("POST")
	r.HandleFunc("/users/{id}", h.IsAuthenticated(controller.Update)).Methods("PATCH")
	r.HandleFunc("/users/{id}", h.IsAuthenticated(controller.Delete)).Methods("DELETE")
}
