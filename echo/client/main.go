package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/178inaba/grpc-example/echo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("k1", "v1-1", "k1", "v1-2", "k2", "v2"))
	r, err := client.Echo(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Fatalf("Did not echo: code: %d, code name: %s, message: %s.", st.Code(), st.Code(), st.Message())
		}

		log.Fatalf("Did not echo: %v.", err)
	}

	log.Println(r.GetMessage())
}
