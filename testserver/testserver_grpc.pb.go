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
	BidiStream(ctx context.Context, opts ...grpc.CallOption) (Sample_BidiStreamClient, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Sample_ClientStreamClient, error)
	ServerStream(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (Sample_ServerStreamClient, error)
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

var sampleBidiStreamStreamDesc = &grpc.StreamDesc{
	StreamName:    "BidiStream",
	ServerStreams: true,
	ClientStreams: true,
}

func (c *sampleClient) BidiStream(ctx context.Context, opts ...grpc.CallOption) (Sample_BidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, sampleBidiStreamStreamDesc, "/Sample/BidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleBidiStreamClient{stream}
	return x, nil
}

type Sample_BidiStreamClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type sampleBidiStreamClient struct {
	grpc.ClientStream
}

func (x *sampleBidiStreamClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleBidiStreamClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var sampleClientStreamStreamDesc = &grpc.StreamDesc{
	StreamName:    "ClientStream",
	ClientStreams: true,
}

func (c *sampleClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Sample_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, sampleClientStreamStreamDesc, "/Sample/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleClientStreamClient{stream}
	return x, nil
}

type Sample_ClientStreamClient interface {
	Send(*SimpleRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type sampleClientStreamClient struct {
	grpc.ClientStream
}

func (x *sampleClientStreamClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleClientStreamClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var sampleServerStreamStreamDesc = &grpc.StreamDesc{
	StreamName:    "ServerStream",
	ServerStreams: true,
}

func (c *sampleClient) ServerStream(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (Sample_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, sampleServerStreamStreamDesc, "/Sample/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sample_ServerStreamClient interface {
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type sampleServerStreamClient struct {
	grpc.ClientStream
}

func (x *sampleServerStreamClient) Recv() (*SimpleResponse, error) {
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
	Simple       func(context.Context, *SimpleRequest) (*SimpleResponse, error)
	BidiStream   func(Sample_BidiStreamServer) error
	ClientStream func(Sample_ClientStreamServer) error
	ServerStream func(*SimpleRequest, Sample_ServerStreamServer) error
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
func (s *SampleService) bidiStream(_ interface{}, stream grpc.ServerStream) error {
	return s.BidiStream(&sampleBidiStreamServer{stream})
}
func (s *SampleService) clientStream(_ interface{}, stream grpc.ServerStream) error {
	return s.ClientStream(&sampleClientStreamServer{stream})
}
func (s *SampleService) serverStream(_ interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.ServerStream(m, &sampleServerStreamServer{stream})
}

type Sample_BidiStreamServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type sampleBidiStreamServer struct {
	grpc.ServerStream
}

func (x *sampleBidiStreamServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleBidiStreamServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type Sample_ClientStreamServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type sampleClientStreamServer struct {
	grpc.ServerStream
}

func (x *sampleClientStreamServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleClientStreamServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type Sample_ServerStreamServer interface {
	Send(*SimpleResponse) error
	grpc.ServerStream
}

type sampleServerStreamServer struct {
	grpc.ServerStream
}

func (x *sampleServerStreamServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

// RegisterSampleService registers a service implementation with a gRPC server.
func RegisterSampleService(s grpc.ServiceRegistrar, srv *SampleService) {
	srvCopy := *srv
	if srvCopy.Simple == nil {
		srvCopy.Simple = func(context.Context, *SimpleRequest) (*SimpleResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Simple not implemented")
		}
	}
	if srvCopy.BidiStream == nil {
		srvCopy.BidiStream = func(Sample_BidiStreamServer) error {
			return status.Errorf(codes.Unimplemented, "method BidiStream not implemented")
		}
	}
	if srvCopy.ClientStream == nil {
		srvCopy.ClientStream = func(Sample_ClientStreamServer) error {
			return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
		}
	}
	if srvCopy.ServerStream == nil {
		srvCopy.ServerStream = func(*SimpleRequest, Sample_ServerStreamServer) error {
			return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
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
				StreamName:    "BidiStream",
				Handler:       srvCopy.bidiStream,
				ServerStreams: true,
				ClientStreams: true,
			},
			{
				StreamName:    "ClientStream",
				Handler:       srvCopy.clientStream,
				ClientStreams: true,
			},
			{
				StreamName:    "ServerStream",
				Handler:       srvCopy.serverStream,
				ServerStreams: true,
			},
		},
		Metadata: "testserver.proto",
	}

	s.RegisterService(&sd, nil)
}
