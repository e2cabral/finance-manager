package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
