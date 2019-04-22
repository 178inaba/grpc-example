package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/178inaba/grpc-example/echo/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v.", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msg := os.Args[1]

	r, err := client.Echo(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		log.Fatalf("Did not echo: %v.", err)
	}

	log.Println(r.GetMessage())
}
