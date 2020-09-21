package grproxy

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc/metadata"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"

	"google.golang.org/grpc"
)

type Proxy struct {
	targetAddr string

	grpc *grpc.Server

	client     *grpc.ClientConn
	reflClient *grpcreflect.Client
}

func NewInterceptor(targetAddr string) *Proxy {
	i := Proxy{
		targetAddr: targetAddr,
	}
	i.grpc = grpc.NewServer(
		grpc.UnaryInterceptor(i.UnaryServerInterceptor()),
	)

	return &i
}

func (p *Proxy) Serve(ctx context.Context, l net.Listener) error {
	errw := func(format string, a ...interface{}) error {
		return fmt.Errorf("grproxy: %w", a...)
	}

	// create a stub server
	_, reflClient, err := p.getClient(ctx)
	if err != nil {
		return errw("get reflection client: %w", err)
	}

	services, err := reflClient.ListServices()
	if err != nil {
		return fmt.Errorf("list services: %w", err)
	}

	for _, svc := range services {
		file, err := reflClient.FileContainingSymbol(svc)
		if err != nil {
			return fmt.Errorf("grproxy: get descriptor: %w", err)
		}
		dsc := file.FindSymbol(svc)
		sd, ok := dsc.(*desc.ServiceDescriptor)
		if !ok {
			return fmt.Errorf("grproxy: target server does not expose service %q", svc)
		}

		grpcSD := grpc.ServiceDesc{
			ServiceName: sd.GetName(),
			Methods:     []grpc.MethodDesc{},
			Streams:     []grpc.StreamDesc{},
			Metadata:    sd.GetFile().GetName(),
		}

		for _, md := range sd.GetMethods() {
			grpcSD.Methods = append(grpcSD.Methods, grpc.MethodDesc{
				MethodName: md.GetName(),
				Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
					msgFactory := dynamic.NewMessageFactoryWithDefaults()
					in := msgFactory.NewDynamicMessage(md.GetInputType())
					if err := dec(in); err != nil {
						return nil, err
					}
					info := &grpc.UnaryServerInfo{
						Server:     nil,
						FullMethod: md.GetFullyQualifiedName(),
					}
					return interceptor(ctx, in, info, nil)
				},
			})
		}

		p.grpc.RegisterService(&grpcSD, nil)
	}

	return p.grpc.Serve(l)
}

func (p *Proxy) UnaryServerInterceptor() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		raw, err := json.MarshalIndent(req, "", "  ")
		if err != nil {
			fmt.Printf("%s\n%v\n", info.FullMethod, err)
			return nil, err
		}
		fmt.Printf("%s\n%s\n", info.FullMethod, raw)

		// proxy call
		client, reflClient, err := p.getClient(ctx)
		if err != nil {
			return nil, err
		}

		// resolve method by name
		svc, mth := parseSymbol(info.FullMethod)
		fullMethodName := svc + "." + mth
		file, err := reflClient.FileContainingSymbol(fullMethodName)
		if err != nil {
			return nil, fmt.Errorf("grproxy: get descriptor: %w", err)
		}
		dsc := file.FindSymbol(svc)
		sd, ok := dsc.(*desc.ServiceDescriptor)
		if !ok {
			return nil, fmt.Errorf("grproxy: target server does not expose service %q", svc)
		}
		mtd := sd.FindMethodByName(mth)
		if mtd == nil {
			return nil, fmt.Errorf("service %q does not include a method named %q", svc, mth)
		}

		// make stub
		msgFactory := dynamic.NewMessageFactoryWithDefaults()
		//req := msgFactory.NewMessage(mtd.GetInputType())
		md := make(metadata.MD)
		ctx = metadata.NewOutgoingContext(ctx, md)
		stub := grpcdynamic.NewStubWithMessageFactory(client, msgFactory)

		msg, ok := req.(proto.Message)
		if !ok {
			return nil, fmt.Errorf("grproxy: req is not an instance of 'proto.Message'")
		}
		resp, err := stub.InvokeRpc(ctx, mtd, msg)
		if err != nil {
			return nil, err
		}

		raw, err = json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Printf("\nresponse:\n%v\n", err)
			return nil, err
		}
		fmt.Printf("\nresponse:\n%s\n", raw)

		return resp, nil
	}
}
func (p *Proxy) StreamServerInterceptor() func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return nil
	}
}

func (p *Proxy) getClient(ctx context.Context) (*grpc.ClientConn, *grpcreflect.Client, error) {
	if p.client == nil {
		cc, err := grpc.DialContext(ctx, p.targetAddr, grpc.WithBlock(), grpc.WithInsecure())
		if err != nil {
			return nil, nil, fmt.Errorf("dial: %w", err)
		}
		p.client = cc
	}
	if p.reflClient == nil {
		p.reflClient = grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(p.client))
	}

	return p.client, p.reflClient, nil
}

func parseSymbol(svcAndMethod string) (string, string) {
	pos := strings.LastIndex(svcAndMethod, "/")
	if pos < 0 {
		pos = strings.LastIndex(svcAndMethod, ".")
		if pos < 0 {
			return "", ""
		}
	}
	svc := svcAndMethod[:pos]
	svc = strings.TrimPrefix(svc, "/")
	return svc, svcAndMethod[pos+1:]
}
