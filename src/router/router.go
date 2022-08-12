package router

import (
	"github.com/MydroX/technical-test-fleurs-d-ici/src/handlers"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	// User
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")

	return r
}
