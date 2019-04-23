package main

import (
	"context"

	pb "github.com/178inaba/grpc-example/echo/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type echoService struct{}

func (s *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	msg := req.GetMessage()
	if msg == "" {
		return nil, status.Error(codes.InvalidArgument, "message is empty")
	}

	return &pb.EchoResponse{Message: msg}, nil
}
