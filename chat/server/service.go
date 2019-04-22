package main

import (
	"io"
	"log"
	"sync"

	pb "github.com/178inaba/grpc-example/chat/proto"
)

var streams sync.Map

type chatService struct{}

func (s *chatService) Connect(stream pb.ChatService_ConnectServer) error {
	log.Printf("Connect: %v.", &stream)
	streams.Store(stream, struct{}{})
	defer func() {
		log.Printf("Disconnect: %v.", &stream)
		streams.Delete(stream)
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		streams.Range(func(key, value interface{}) bool {
			stream := key.(pb.ChatService_ConnectServer)
			stream.Send(&pb.Post{
				Name:    req.GetName(),
				Message: req.GetMessage(),
			})

			return true
		})
	}
}
