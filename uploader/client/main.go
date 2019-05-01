package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/178inaba/grpc-example/interceptor"
	pb "github.com/178inaba/grpc-example/uploader/proto"
	"google.golang.org/grpc"
)

func main() {
	name := os.Args[1]
	fs, err := os.Open(name)
	if err != nil {
		log.Fatalf("Could not open file: %v.", err)
	}
	defer fs.Close()

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.LoggingUnaryForClient),
		grpc.WithStreamInterceptor(interceptor.LoggingStreamForClient))
	if err != nil {
		log.Fatalf("Did not connect: %v.", err)
	}
	defer conn.Close()

	c := pb.NewFileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := c.Upload(ctx)
	if err != nil {
		log.Fatalf("Could not upload file: %v.", err)
	}

	buf := make([]byte, 1000*1024)
	for {
		n, err := fs.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Could not read file: %v.", err)
		}

		if err := stream.Send(&pb.FileRequest{Name: filepath.Base(name), Data: buf[:n]}); err != nil {
			log.Printf("Could not send file: %v.", err)
			break
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not receive response file: %v.", err)
	}

	log.Printf("Done %d bytes.", res.GetSize())
}
