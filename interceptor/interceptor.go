package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggingUnaryForClient(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("Method: %s, ElapsedTime: %s.", method, time.Since(start))

	return err
}

func LoggingUnaryForServer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	h, err := handler(ctx, req)
	log.Printf("Method: %s, ElapsedTime: %s, StatusCode: %s", info.FullMethod, time.Since(start), status.Code(err))

	return h, err
}
