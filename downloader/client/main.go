package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/178inaba/grpc-example/downloader/proto"
	"github.com/178inaba/grpc-example/interceptor"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.LoggingUnaryForClient))
	if err != nil {
		log.Fatalf("Did not connect: %v.", err)
	}

	c := pb.NewFileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	name := os.Args[1]

	stream, err := c.Download(ctx, &pb.FileRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not download: %v.", err)
	}

	var blob []byte
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Done %d bytes.", len(blob))
			break
		} else if err != nil {
			log.Fatalf("Could not receive: %v.", err)
		}

		blob = append(blob, c.GetData()...)
	}

	ioutil.WriteFile(name, blob, 0644)
}
