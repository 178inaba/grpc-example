package main

import (
	"io"
	"io/ioutil"
	"log"
	"path/filepath"

	pb "github.com/178inaba/grpc-example/uploader/proto"
)

type fileService struct{}

func (s *fileService) Upload(stream pb.FileService_UploadServer) error {
	var name string
	var blob []byte
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Done %d bytes.", len(blob))
			break
		} else if err != nil {
			return err
		}

		name = c.GetName()
		blob = append(blob, c.GetData()...)
	}

	fp := filepath.Join("./uploader/resource", name)
	if err := ioutil.WriteFile(fp, blob, 0644); err != nil {
		return err
	}

	return stream.SendAndClose(&pb.FileResponse{Size: int64(len(blob))})
}
