package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type healthHandler struct{}

// NewHealthHandler creates a new health check handler
func NewHealthHandler() http.Handler {
	return &healthHandler{}
}

func (h *healthHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	logger := log.WithFields(
		log.Fields{
			"method":     req.Method,
			"path":       req.URL.Path,
			"user-agent": req.UserAgent()})

	if req.Method != "GET" {
		wr.WriteHeader(400)
		logger.Error("Bad method")
		return
	}

	switch req.URL.Path {
	case "/health/live":
		fallthrough
	case "/health/ready":
		wr.WriteHeader(200)
		logger.Info("Server is live and ready")
		return
	default:
		wr.WriteHeader(404)
		logger.Error("Bad path")
	}
}
