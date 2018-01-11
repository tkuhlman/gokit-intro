package main

import (
	"net/http"
	"os"

	"github.com/tkuhlman/gokit-intro/service"

	"github.com/go-kit/kit/ratelimit"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

func main() {
	// init dependencies
	log := logrus.New()

	// A new Content service
	svc := service.NewV1()
	// Add service level middleware
	svc = service.LogMiddleware(log)(svc)

	// Build endpoints then add endpoint middleware
	queryEndpoint := makeQueryEndpoint(svc)
	queryEndpoint = ratelimit.NewDelayingLimiter(rate.NewLimiter(50, 10))(queryEndpoint)
	queryEndpoint = newRelicMiddleware(log, fauxNewRelic{name: "query"})(queryEndpoint)

	// Use the Go kit included http transport layer with a new Query Endpoint to create a handler
	queryHandler := httptransport.NewServer(queryEndpoint, decodeRequest, encodeResponse)

	// Enable the handler and serve traffic
	http.Handle("/v1", queryHandler)
	log.Info("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
