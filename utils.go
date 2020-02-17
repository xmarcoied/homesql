package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// DefaultGetenv tries to getenv and returns a default value
func DefaultGetenv(k, d string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Printf("[%s] environment variable is not set; ==> setting [%s] to default value = %s", k, k, d)
		return d
	}
	return v
}

// ReturnSuccessMessage encodes the passed msg to json,
// then returns a json response with status_code 200
func ReturnSuccessMessage(w http.ResponseWriter, msg interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

// ReturnErrorMessage encodes the passed msg to json,
// then returns a json error response with status_code 500
func ReturnErrorMessage(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
