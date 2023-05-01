package grpc_server

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

//TODO: apply interceptor to gRPC server handler

func MetricsUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	// pre-processing logic
	log.Printf("Received request for method %s", info.FullMethod)

	// call the next handler in the chain
	resp, err := handler(ctx, req)
	if err != nil {
		// error handling logic
		return nil, err
	}
	// post-processing logic
	return resp, err
}
