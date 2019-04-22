package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/178inaba/grpc-example/chat/proto"
	"google.golang.org/grpc"
)

func main() {
	name := os.Args[1]

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v.", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)
	stream, err := c.Connect(context.Background())
	if err != nil {
		log.Fatalf("Could not connect: %v.", err)
	}

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				return
			} else if err != nil {
				log.Fatalf("Failed to recv: %v.", err)
			}

			fmt.Printf("[%s] %s\n", res.GetName(), res.GetMessage())
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg := scanner.Text()
		if msg == ":quit" {
			stream.CloseSend()
			return
		}

		stream.Send(&pb.Post{Name: name, Message: msg})
	}
}
