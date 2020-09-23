// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package testserver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SampleClient is the client API for Sample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleClient interface {
	Simple(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	Streaming(ctx context.Context, opts ...grpc.CallOption) (Sample_StreamingClient, error)
}

type sampleClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleClient(cc grpc.ClientConnInterface) SampleClient {
	return &sampleClient{cc}
}

var sampleSimpleStreamDesc = &grpc.StreamDesc{
	StreamName: "Simple",
}

func (c *sampleClient) Simple(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/Sample/Simple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sampleStreamingStreamDesc = &grpc.StreamDesc{
	StreamName:    "Streaming",
	ServerStreams: true,
	ClientStreams: true,
}

func (c *sampleClient) Streaming(ctx context.Context, opts ...grpc.CallOption) (Sample_StreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, sampleStreamingStreamDesc, "/Sample/Streaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleStreamingClient{stream}
	return x, nil
}

type Sample_StreamingClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type sampleStreamingClient struct {
	grpc.ClientStream
}

func (x *sampleStreamingClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleStreamingClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SampleService is the service API for Sample service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSampleService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SampleService struct {
	Simple    func(context.Context, *SimpleRequest) (*SimpleResponse, error)
	Streaming func(Sample_StreamingServer) error
}

func (s *SampleService) simple(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Simple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/Sample/Simple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Simple(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SampleService) streaming(_ interface{}, stream grpc.ServerStream) error {
	return s.Streaming(&sampleStreamingServer{stream})
}

type Sample_StreamingServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type sampleStreamingServer struct {
	grpc.ServerStream
}

func (x *sampleStreamingServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleStreamingServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterSampleService registers a service implementation with a gRPC server.
func RegisterSampleService(s grpc.ServiceRegistrar, srv *SampleService) {
	srvCopy := *srv
	if srvCopy.Simple == nil {
		srvCopy.Simple = func(context.Context, *SimpleRequest) (*SimpleResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Simple not implemented")
		}
	}
	if srvCopy.Streaming == nil {
		srvCopy.Streaming = func(Sample_StreamingServer) error {
			return status.Errorf(codes.Unimplemented, "method Streaming not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "Sample",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Simple",
				Handler:    srvCopy.simple,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Streaming",
				Handler:       srvCopy.streaming,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "testserver.proto",
	}

	s.RegisterService(&sd, nil)
}
