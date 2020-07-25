// Package router is responsible for setting up all the endpoints supported by the system.
package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HealthController Object providing over all status of the service
type HealthController interface {
	Status(w http.ResponseWriter, r *http.Request)
}

// Init initializes the endpoints we will handle
func Init(h HealthController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/status", h.Status).Methods("GET")
	return r
}
