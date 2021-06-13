package grpc

import (
	"context"
	"errors"

	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/technolingo/gokittodo/users/pkg/endpoint"
	pb "github.com/technolingo/gokittodo/users/pkg/grpc/pb"
)

func makeCreateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)
}

// decodeCreateRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'users' decoder is not impelemented")
}

// encodeCreateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'users' encoder is not impelemented")
}

// Create
func (c *usersClient) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	rep, err := c.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateReply), nil
}
