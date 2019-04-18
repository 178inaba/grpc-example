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
	target := "localhost:50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msg := os.Args[1]

	r, err := client.Echo(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		log.Println(err)
	}

	log.Println(r.GetMessage())
}
