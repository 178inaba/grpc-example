package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// LoggingUnaryForClient logs Unary RPC client side.
func LoggingUnaryForClient(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("Unary - Method: %s, StatusCode: %s, ElapsedTime: %s.", method, status.Code(err), time.Since(start))

	return err
}

// LoggingUnaryForServer logs Unary RPC server side.
func LoggingUnaryForServer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	h, err := handler(ctx, req)
	log.Printf("Unary - Method: %s, StatusCode: %s, ElapsedTime: %s.", info.FullMethod, status.Code(err), time.Since(start))

	return h, err
}

// LoggingStreamForClient performs streaming RPC client-side logging.
func LoggingStreamForClient(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	stream, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("Stream - Method: %s, StatusCode: %s.", method, status.Code(err))

	return stream, err
}

// LoggingStreamForServer performs streaming RPC server-side logging.
func LoggingStreamForServer(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	err := handler(srv, ss)
	log.Printf("Stream - Method: %s, StatusCode: %s, ConnectionTime: %s.", info.FullMethod, status.Code(err), time.Since(start))

	return err
}
