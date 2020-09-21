package testserver

import (
	"context"
)

type sampleHandler struct{}

func NewHandler() *SampleService {
	var h sampleHandler
	return &SampleService{
		Simple: h.Simple,
	}
}

func (s sampleHandler) Simple(_ context.Context, req *SimpleRequest) (*SimpleResponse, error) {
	return &SimpleResponse{
		Attr1: "response " + req.Attr1,
	}, nil
}
