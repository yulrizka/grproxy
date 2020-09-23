package testserver

import (
	"context"
	"fmt"
	"io"
	"strings"
)

type sampleHandler struct{}

func NewHandler() *SampleService {
	var h sampleHandler
	return &SampleService{
		Simple:       h.Simple,
		Streaming:    h.Streaming,
		ClientStream: h.ClientStream,
		ServerStream: h.ServerStream,
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

func (s sampleHandler) ClientStream(ss Sample_ClientStreamServer) error {
	var out strings.Builder
	i := 0
	for {
		msg, err := ss.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("recv: %w", err)
		}
		out.WriteString(fmt.Sprintf("%d received: %s;", i, msg.Attr1))
		i++
	}

	return ss.SendAndClose(&SimpleResponse{Attr1: out.String()})
}

func (s sampleHandler) ServerStream(req *SimpleRequest, ss Sample_ServerStreamServer) error {
	for i := 0; i < 2; i++ {
		resp := SimpleResponse{Attr1: fmt.Sprintf("received %s %d", req.Attr1, i)}
		if err := ss.Send(&resp); err != nil {
			return fmt.Errorf("send reponse: %w", err)
		}
	}

	return nil
}
