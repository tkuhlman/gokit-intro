package main

import (
	"context"
	"errors"

	"github.com/tkuhlman/gokit-intro/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
)

// Endpoints follow an RPC style with standard structs for the request and response.

type queryRequest struct {
	query string
}

type queryResponse struct {
	resp string
}

// makeQueryEndpoint wraps a call to the content service in an endpoint style with the appropriate RPC response/request
// structs. This simple example has a 1 to 1 mapping between the endpoint and service call but a single endpoint could
// result in multiple calls to the service in a more complicated service.
func makeQueryEndpoint(svc service.Content) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(queryRequest)
		if !ok {
			return nil, errors.New("failed request type assertion")
		}
		resp, err := svc.Query(ctx, req.query)
		if err != nil {
			return nil, err
		}
		return queryResponse{resp}, nil
	}
}

// fauxNewRelic is a fake New Relic connection used for example middleware.
type fauxNewRelic struct {
	name string
}

func newRelicMiddleware(log *logrus.Logger, nr fauxNewRelic) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			log.Infof("Starting new relic transaction for %q", nr.name)
			defer log.Infof("Finishing new relic transaction for %q", nr.name)
			return next(ctx, request)
		}
	}
}
