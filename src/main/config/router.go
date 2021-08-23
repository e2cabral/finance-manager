package config

import (
	routes2 "finance-manager/src/main/routes"
	"github.com/gorilla/mux"
)

type RoutesConfig struct {
	Prefix string
}

func NewRoutesConfig(prefix string) *RoutesConfig {
	return &RoutesConfig{Prefix: prefix}
}

func (c *RoutesConfig) SetupRoutes(r *mux.Router) {
	sub := r.PathPrefix(c.Prefix).Subrouter()
	routes2.MovementsRoutes(sub)
	routes2.PocketRoutes(sub)
	routes2.UsersRoutes(sub)
	routes2.AuthRoutes(sub)
}
