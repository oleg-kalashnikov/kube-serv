package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oleg-kalashnikov/kube-serv/pkg/config"
	"github.com/oleg-kalashnikov/kube-serv/pkg/handlers"
	"github.com/oleg-kalashnikov/kube-serv/pkg/logger"
	stdlog "github.com/oleg-kalashnikov/kube-serv/pkg/logger/standard"
	"github.com/oleg-kalashnikov/kube-serv/pkg/version"
)

// Setup configures the service
func Setup(c *config.Config) (r *mux.Router, log logger.Logger, err error) {
	// Setup logger
	log = stdlog.New(&logger.Config{
		Level: c.LogLevel,
		Time:  true,
		UTC:   true,
	})

	log.Info("Version:", version.RELEASE)
	log.Warnf("%s log level is used", logger.LevelDebug.String())
	log.Infof("Service %s listened on %s:%d", config.SERVICENAME, c.LocalHost, c.LocalPort)

	// Define handlers
	h := handlers.New(log, c)

	// Register new router
	r = mux.NewRouter()

	// Response for undefined methods
	r.NotFoundHandler = http.HandlerFunc(notFound)

	// Configure router
	r.Use(h.CaptureMetrics)
	r.HandleFunc("/", h.Root).Methods("GET")
	r.HandleFunc("/healthz", h.Health).Methods("GET")
	r.HandleFunc("/readyz", h.Ready).Methods("GET")
	r.HandleFunc("/info", h.Info).Methods("GET")

	return r, log, err
}

// Response for undefined methods
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}
