package main

import (
	"finance-manager/domain/models"
	database "finance-manager/infra/database/client"
	"finance-manager/main/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	m := mux.NewRouter()

	router := config.NewRoutesConfig("/v1/api/")
	router.SetupRoutes(m)

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&models.Movement{}); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", m))
}
