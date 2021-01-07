package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/yulrizka/grproxy"
)

func main() {
	grpcAddr := ":9999"
	httpAddr := ":9998"
	target := "127.0.0.1:10000"

	flag.StringVar(&grpcAddr, "grpc-addr", env("GRPC_ADDR", grpcAddr), "proxy address")
	flag.StringVar(&httpAddr, "http-addr", env("GRPC_ADDR", httpAddr), "proxy address")
	flag.StringVar(&target, "target", env("TARGET", target), "target grpc server")
	flag.Parse()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	log.Printf("grproxy %v -> %v", grpcAddr, target)
	log.Printf("HTTP UI %v", httpAddr)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		interceptor := grproxy.NewInterceptor(target)
		return interceptor.Serve(ctx, l)
	})

	g.Go(func() error {
		httpHandler := grproxy.NewHTTPHandler()
		mux := http.NewServeMux()
		mux.HandleFunc("/ws", httpHandler.HandleWebsocket)

		httpServer := &http.Server{
			Addr:              httpAddr,
			ReadHeaderTimeout: 10 * time.Second,
			WriteTimeout:      10 * time.Second,
			Handler:           mux,
		}

		go func() {
			<-ctx.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = httpServer.Shutdown(ctx)
		}()

		return httpServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatalf(err.Error())
	}
}

func env(s string, def string) string {
	v := os.Getenv(s)
	if v == "" {
		return def
	}
	return v
}
