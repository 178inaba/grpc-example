package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func LoggingUnaryForClient(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("Unary - Method: %s, ElapsedTime: %s.", method, time.Since(start))

	return err
}
