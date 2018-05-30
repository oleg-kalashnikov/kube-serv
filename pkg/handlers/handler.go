package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/oleg-kalashnikov/kube-serv/pkg/config"
	"github.com/oleg-kalashnikov/kube-serv/pkg/logger"
	"github.com/oleg-kalashnikov/kube-serv/pkg/version"
)

// Handler defines common part for all handlers
type Handler struct {
	logger      logger.Logger
	config      *config.Config
	maintenance bool
	stats       *stats
}

type stats struct {
	requests        *Requests
	averageDuration time.Duration
	maxDuration     time.Duration
	totalDuration   time.Duration
	requestsCount   time.Duration
	startTime       time.Time
}

// New returns new instance of the Handler
func New(logger logger.Logger, config *config.Config) *Handler {
	return &Handler{
		logger: logger,
		config: config,
		stats: &stats{
			requests:  new(Requests),
			startTime: time.Now(),
		},
	}
}

// Root handler shows version
func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf(`{"name": %s, "version": %s}`, config.SERVICENAME, version.RELEASE))
	// json.NewEncoder(w).Encode(status)
}
