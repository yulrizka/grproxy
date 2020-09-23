package testserver

import (
	"context"
	"io"
)

type sampleHandler struct{}

func NewHandler() *SampleService {
	var h sampleHandler
	return &SampleService{
		Simple:    h.Simple,
		Streaming: h.Streaming,
	}
}

func (s sampleHandler) Simple(_ context.Context, req *SimpleRequest) (*SimpleResponse, error) {
	return &SimpleResponse{
		Attr1: "response " + req.Attr1,
	}, nil
}

func (s sampleHandler) Streaming(stream Sample_StreamingServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err = stream.Send(&SimpleResponse{Attr1: "response " + in.Attr1}); err != nil {
			return err
		}
	}
}
