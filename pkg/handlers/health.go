package handlers

import (
	"io"
	"net/http"
)

// Health returns "OK" if service is alive
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, http.StatusText(http.StatusOK))
}
