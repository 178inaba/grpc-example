package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/178inaba/grpc-example/downloader/proto"
	"google.golang.org/grpc"
)

func main() {
	target := "localhost:50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s.", err)
	}

	c := pb.NewFileServiceClient(conn)
	name := os.Args[1]
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := c.Download(ctx, &pb.FileRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not download: %s\n", err)
	}

	var blob []byte
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Done %d bytes\n", len(blob))
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		blob = append(blob, c.GetData()...)
	}

	ioutil.WriteFile(name, blob, 0644)
}
