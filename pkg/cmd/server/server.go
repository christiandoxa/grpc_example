package cmd

import (
	"context"
	"grpc_bri/pkg/protocol/grpc"
	"grpc_bri/pkg/repository/v1"
	"grpc_bri/pkg/service/v1" //nolint:typecheck
	"os"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	v1API := v1.NewToDoServiceServer(repository.New())

	return grpc.RunServer(ctx, v1API, os.Getenv("GRPC_PORT"))
}
