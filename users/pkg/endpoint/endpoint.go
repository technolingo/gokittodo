package endpoint

import (
	"context"
	service "users/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		e0 := s.Create(ctx, req.Email, req.Password)
		return CreateResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.E0
}

// RetrieveRequest collects the request parameters for the Retrieve method.
type RetrieveRequest struct {
	Email string `json:"email"`
}

// RetrieveResponse collects the response parameters for the Retrieve method.
type RetrieveResponse struct {
	E0 error `json:"e0"`
}

// MakeRetrieveEndpoint returns an endpoint that invokes Retrieve on the service.
func MakeRetrieveEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RetrieveRequest)
		e0 := s.Retrieve(ctx, req.Email)
		return RetrieveResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r RetrieveResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, email string, password string) (e0 error) {
	request := CreateRequest{
		Email:    email,
		Password: password,
	}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).E0
}

// Retrieve implements Service. Primarily useful in a client.
func (e Endpoints) Retrieve(ctx context.Context, email string) (e0 error) {
	request := RetrieveRequest{Email: email}
	response, err := e.RetrieveEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RetrieveResponse).E0
}
