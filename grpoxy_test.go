package grproxy

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/yulrizka/grproxy/testserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func TestSimpleCall(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	target, err := newServer()
	if err != nil {
		t.Fatal(err)
	}
	if err := target.Start(context.Background()); err != nil {
		t.Fatal(err)
	}

	// start the proxy
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			t.Fatal(err)
		}
	}
	interceptor := NewInterceptor(target.listener.Addr().String())
	go func() {
		if err := interceptor.Serve(ctx, l); err != nil {
			t.Logf("interceptor server: %v", err)
		}
	}()

	// test calling function
	addr := l.Addr().String()
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatal(err)
	}
	client := testserver.NewSampleClient(conn)

	req := testserver.SimpleRequest{Attr1: "attr1"}
	resp, err := client.Simple(ctx, &req)
	if err != nil {
		t.Fatalf("simple: %v", err)
	}

	if got, want := resp.Attr1, "response "+req.Attr1; got != want {
		t.Fatalf("attr1 got %q want %q", got, want)
	}

}

type server struct {
	grpc     *grpc.Server
	listener net.Listener
}

func newServer() (*server, error) {
	s := new(server)
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			return nil, fmt.Errorf("httptest: failed to listen on a port: %v", err)
		}
	}
	s.listener = l

	s.grpc = grpc.NewServer()
	testserver.RegisterSampleService(s.grpc, testserver.NewHandler())
	reflection.Register(s.grpc)

	return s, nil
}

func (s *server) Start(ctx context.Context) error {
	go func() {
		if err := s.grpc.Serve(s.listener); err != nil {
			panic(fmt.Sprintf("grpc server: %v", err))
		}
	}()
	go func() {
		<-ctx.Done()
		s.grpc.Stop()

	}()

	return nil
}
