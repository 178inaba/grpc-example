package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"

	pb "github.com/178inaba/grpc-example/echo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type echoService struct{}

func (s *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	msg := req.GetMessage()
	if msg == "" {
		return nil, status.Error(codes.InvalidArgument, "message is empty")
	}

	log.Printf("Message is %q.", msg)

	h := sha256.New()
	h.Write([]byte(msg))
	hash := fmt.Sprintf("%x", h.Sum(nil))

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Metadata is %v.", md)

		clientHash := md.Get("sha256")[0]
		log.Printf("client message hash is %q.", clientHash)
		log.Printf("server message hash is %q.", hash)
		if clientHash != hash {
			return nil, status.Error(codes.InvalidArgument, "the message has been tampered with")
		}

		grpc.SendHeader(ctx, md)
	} else {
		return nil, status.Error(codes.InvalidArgument, "sha256 not found in header")
	}

	grpc.SetTrailer(ctx, metadata.Pairs("sha256", hash))
	return &pb.EchoResponse{Message: msg}, nil
}
