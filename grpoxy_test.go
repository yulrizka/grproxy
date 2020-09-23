package grproxy

import (
	"context"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/yulrizka/grproxy/testserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func TestCall(t *testing.T) {
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

	// test calling unary function
	addr := l.Addr().String()
	//addr = target.listener.Addr().String()
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatal(err)
	}
	client := testserver.NewSampleClient(conn)

	t.Run("unary", func(t *testing.T) {
		req := testserver.SimpleRequest{Attr1: "attr1"}
		resp, err := client.Simple(ctx, &req)
		if err != nil {
			t.Fatalf("simple: %v", err)
		}

		if got, want := resp.Attr1, "response "+req.Attr1; got != want {
			t.Fatalf("attr1 got %q want %q", got, want)
		}
	})

	t.Run("stream", func(t *testing.T) {
		stream, err := client.BidiStream(ctx)
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 2; i++ {
			req := testserver.SimpleRequest{Attr1: fmt.Sprintf("attr1 %d", i)}
			err := stream.Send(&req)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := stream.Recv()
			if err != nil {
				t.Fatal(err)
			}
			if got, want := resp.Attr1, "response "+req.Attr1; got != want {
				t.Errorf("got %q want %q", got, want)
			}
		}
		if err := stream.CloseSend(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("client stream", func(t *testing.T) {
		stream, err := client.ClientStream(ctx)
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 2; i++ {
			req := testserver.SimpleRequest{Attr1: fmt.Sprintf("attr1 %d", i)}
			err := stream.Send(&req)
			if err != nil {
				t.Fatal(err)
			}
		}
		if err := stream.CloseSend(); err != nil {
			t.Fatal(err)
		}
		var resp testserver.SimpleResponse
		if err = stream.RecvMsg(&resp); err != nil {
			t.Fatal(err)
		}
		got, want := resp.Attr1, "0 received: attr1 0;1 received: attr1 1;"
		if got != want {
			t.Fatalf("got %v want %v", got, want)
		}
	})

	t.Run("server stream", func(t *testing.T) {
		stream, err := client.ServerStream(ctx, &testserver.SimpleRequest{Attr1: fmt.Sprintf("attr1")})
		if err != nil {
			t.Fatal(err)
		}

		// receive multiple message
		var s strings.Builder
		for i := 0; i < 2; i++ {
			resp, err := stream.Recv()
			if err != nil {
				t.Fatal(err)
			}
			s.WriteString(resp.Attr1)
			s.WriteString(";")
		}
		if got, want := s.String(), "received attr1 0;received attr1 1;"; got != want {
			t.Fatalf("got %s want %s", got, want)
		}

		if err := stream.CloseSend(); err != nil {
			t.Fatal(err)
		}

	})
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
