package routes

import (
	h "finance-manager/src/infra/helpers/auth"
	"finance-manager/src/presentation/controllers"
	"github.com/gorilla/mux"
)

func PocketRoutes(r *mux.Router) {
	controller := controllers.NewPocketController()
	r.HandleFunc("/pockets", h.IsAuthenticated(controller.GetPockets)).Methods("GET")
	r.HandleFunc("/pockets/{id}", h.IsAuthenticated(controller.GetPocketById)).Methods("GET")
	r.HandleFunc("/pockets", h.IsAuthenticated(controller.Save)).Methods("POST")
	r.HandleFunc("/pockets/{id}", h.IsAuthenticated(controller.Update)).Methods("PATCH")
	r.HandleFunc("/pockets/{id}", h.IsAuthenticated(controller.Delete)).Methods("DELETE")
}
