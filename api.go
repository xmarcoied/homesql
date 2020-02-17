package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func registerHandlers() {
	// TODO: Checking http://www.gorillatoolkit.org/pkg/mux
	router := mux.NewRouter()

	router.HandleFunc("/api/user/", userHandler).Methods("POST")
	router.HandleFunc("/api/home/", homeHandler).Methods("POST")

	router.HandleFunc("/api/users/", usersHandler).Methods("GET")

	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	// [END request_logging]
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	userRequest := User{}
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		ReturnErrorMessage(w, err)
		return
	}

	userResponse, err := DB.NewUser(userRequest)
	if err != nil {
		ReturnErrorMessage(w, err)
		return
	}

	ReturnSuccessMessage(w, userResponse)
	return
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	homeRequest := Home{}
	err := json.NewDecoder(r.Body).Decode(&homeRequest)
	if err != nil {
		ReturnErrorMessage(w, err)
		return
	}

	homeResponse, err := DB.NewHome(homeRequest)
	if err != nil {
		ReturnErrorMessage(w, err)
		return
	}

	ReturnSuccessMessage(w, homeResponse)
	return
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := DB.GetUsers(r.Context())
	if err != nil {
		ReturnErrorMessage(w, err)
		return
	}

	ReturnSuccessMessage(w, users)
	return
}
