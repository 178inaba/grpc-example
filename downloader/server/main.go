package main

import (
	"log"
	"net"

	pb "github.com/178inaba/grpc-example/downloader/proto"
	"google.golang.org/grpc"
)

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	srv := grpc.NewServer()
	pb.RegisterFileServiceServer(srv, &fileService{})
	log.Printf("Start server on port%s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v\n", err)
	}
}
