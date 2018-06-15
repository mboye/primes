package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mboye/primes"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type primeHandler struct {
	checker  primes.Checker
	counters map[int]prometheus.Counter
}

// NewPrimeHandler created a new HTTP request handler for checking if a number is a prime
func NewPrimeHandler(checker primes.Checker) http.Handler {
	return &primeHandler{
		checker:  checker,
		counters: make(map[int]prometheus.Counter)}
}

func (h *primeHandler) getCounter(statusCode int) prometheus.Counter {
	if counter, exists := h.counters[statusCode]; exists {
		return counter
	}

	newCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name:        "http_requests_total",
			Help:        "Total number of HTTP requests",
			ConstLabels: map[string]string{"status_code": strconv.Itoa(statusCode)}})

	prometheus.Register(newCounter)
	h.counters[statusCode] = newCounter
	return newCounter
}

func (h *primeHandler) incrementCounter(statusCode int) {
	h.getCounter(statusCode).Inc()
}

func (h *primeHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	logger := log.WithFields(
		log.Fields{
			"method":     req.Method,
			"path":       req.URL.Path,
			"user-agent": req.UserAgent()})

	if req.Method != "GET" {
		logger.Warn("Bad HTTP method")
		wr.WriteHeader(400)
		h.incrementCounter(400)
		return
	}

	valueStr := req.FormValue("value")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		logger.Error("Bad value parameter in request")
		wr.WriteHeader(400)
		h.incrementCounter(400)
		return
	}

	isPrime := h.checker.IsPrime(value)

	responseData, err := json.Marshal(
		struct {
			Value   int  `json:"value"`
			IsPrime bool `json:"isPrime"`
		}{value, isPrime})

	if err != nil {
		log.Error("Failed to marshal JSON response")
		wr.WriteHeader(500)
		h.incrementCounter(500)
		return
	}

	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(200)

	_, err = wr.Write(responseData)
	if err != nil {
		logger.Errorf("Failed to write response: %s", err.Error())
	} else {
		logger.WithFields(log.Fields{"value": value, "isPrime": isPrime}).Info("Processed request")
		h.incrementCounter(200)
	}
}
