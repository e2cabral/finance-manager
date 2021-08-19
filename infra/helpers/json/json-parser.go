package helpers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ToJSON(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func FromJSON(r io.ReadCloser, data interface{}) error {
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return err
	}
	return nil
}
