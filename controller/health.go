package controller

import "net/http"

// Health controller for the health of the system
type Health struct{}

// swagger:route GET /status status
// checks if the serivce is reachable
// responses:
//   200: StatusOK

// Status checks if the serivce is running
func (h Health) Status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\nStatus OK for now\n"))
	w.WriteHeader(http.StatusOK)
}

// New Create a health object controller
func New() *Health {
	return &Health{}
}
