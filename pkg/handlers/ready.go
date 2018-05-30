package handlers

import (
	"io"
	"net/http"
)

// Ready returns "OK" if service is ready to serve traffic
func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	// TODO: possible use cases:
	// load data from a database, a message broker, any external services, etc

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, http.StatusText(http.StatusOK))
}
