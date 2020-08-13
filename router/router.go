package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HealthController Object providing over all status of the service
type HealthController interface {
	Status(w http.ResponseWriter, r *http.Request)
}

// EventController handles event related requests
type EventController interface {
	Create(resp http.ResponseWriter, request *http.Request)
}

// Init initializes the endpoints we will handle
func Init(h HealthController, eventController EventController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/status", h.Status).Methods("GET")

	eventRouter := r.PathPrefix("/v1/user/{user_id}")

	eventRouter.PathPrefix("/event").HandlerFunc(eventController.Create).Methods("POST")

	return r
}
