package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/178inaba/grpc-example/echo/proto"
	"github.com/178inaba/grpc-example/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.LoggingUnaryForClient),
		grpc.WithStreamInterceptor(interceptor.LoggingStreamForClient))
	if err != nil {
		log.Fatalf("Did not connect: %v.", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if len(os.Args) < 2 {
		log.Fatal("Please enter a message.")
	}

	msg := os.Args[1]

	h := sha256.New()
	h.Write([]byte(msg))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	log.Printf("Message sha256 is %q.", hash)

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("k1", "v1-1", "k1", "v1-2", "k2", "v2", "sha256", hash))
	var header, trailer metadata.MD
	r, err := client.Echo(ctx, &pb.EchoRequest{Message: msg}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Fatalf("Did not echo: code: %d, code name: %s, message: %s.", st.Code(), st.Code(), st.Message())
		}

		log.Fatalf("Did not echo: %v.", err)
	}

	serverHash := trailer.Get("sha256")[0]
	log.Printf("Client hash is %q.", hash)
	log.Printf("Server hash is %q.", serverHash)
	if serverHash != hash {
		log.Fatal("The message has been tampered with.")
	}

	log.Printf("Server echo header metadata: %v.", header)
	log.Printf("Server echo message is %q.", r.GetMessage())
}
