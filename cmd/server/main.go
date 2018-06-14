package main

import (
	"flag"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mboye/primes/checkers/fast"
	"github.com/mboye/primes/checkers/slow"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/mboye/primes"

	log "github.com/sirupsen/logrus"

	"github.com/mboye/primes/cmd/server/handler"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen for HTTP requests on")
	algorithm := flag.String("algorithm", "slow", "Prime checking algorithm: slow or fast")
	flag.Parse()

	logger := log.New()

	var checker primes.Checker
	switch strings.ToLower(*algorithm) {
	case "slow":
		checker = slow.New()
		logger.Info("Using slow prime checking algorithm")
	case "fast":
		checker = fast.New()
		logger.Info("Using fast prime checking algorithm")
	default:
		logger.Errorf("Invalid prime checking algorithm: %s", *algorithm)
		os.Exit(1)
	}

	go runMetricsServer()

	primeHandler := handler.NewPrimeHandler(checker)
	http.Handle("/api/isPrime", primeHandler)

	logger.WithField("port", *port).Info("Starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.WithField("error", err.Error()).Fatal("Failed to start server")
	}
}

func runMetricsServer() {
	logger := log.New()

	port := 9095
	server := http.NewServeMux()
	server.Handle("/metrics", promhttp.Handler())

	logger.WithField("port", port).Info("Starting metrics server")
	err := http.ListenAndServe(":"+strconv.Itoa(port), server)
	if err != nil {
		logger.WithField("error", err.Error()).Fatal("Failed to start metrics server")
	}
}
