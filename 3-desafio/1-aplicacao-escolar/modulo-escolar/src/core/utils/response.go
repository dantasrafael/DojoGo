package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode((data)); err != nil {
		log.Fatal(err)
	}
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	JsonResponse(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
