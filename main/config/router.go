package config

import (
	"finance-manager/main/routes"
	"github.com/gorilla/mux"
)

type RoutesConfig struct {
	Prefix string
}

func NewRoutesConfig(prefix string) *RoutesConfig {
	return &RoutesConfig{Prefix: prefix}
}

func (c *RoutesConfig) SetupRoutes(r *mux.Router) {
	routes.MovementsRoutes(r.PathPrefix(c.Prefix).Subrouter())
}
