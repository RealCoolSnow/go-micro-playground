package endpoint

import (
	"context"
	"micro-playgound/service"

	"github.com/go-kit/kit/endpoint"
)

type DiscoveryEndpoints struct {
	SayHelloEndpoint    endpoint.Endpoint
	UppercaseEndpoint   endpoint.Endpoint
	DiscoveryEndpoint   endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

// say hello
type SayHelloRequest struct {
}

type SayHelloResponse struct {
	Message string `json:"message"`
}

func MakeSayHelloEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		message := svc.SayHello()
		return SayHelloResponse{
			Message: message,
		}, nil
	}
}

// uppercase
type UppercaseRequest struct {
	Text string
}

type UppercaseResponse struct {
	UpperText string `json:"upper_text"`
}

func MakeUpperEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UppercaseRequest)
		text := svc.Uppercase(req.Text)
		return UppercaseResponse{
			UpperText: text,
		}, nil
	}
}

// discover service
type DiscoveryRequest struct {
	ServiceName string
}

type DiscoveryResponse struct {
	Instances []interface{} `json:"instances"`
	Error     string        `json:"error"`
}

func MakeDiscoveryEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DiscoveryRequest)
		instances, err := svc.DiscoveryService(ctx, req.ServiceName)
		var errString = ""
		if err != nil {
			errString = err.Error()
		}
		return &DiscoveryResponse{
			Instances: instances,
			Error:     errString,
		}, nil
	}
}

// health check
type HealthRequest struct{}

type HealthResponse struct {
	Status bool `json:"status"`
}

func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()
		return HealthResponse{
			Status: status,
		}, nil
	}
}
