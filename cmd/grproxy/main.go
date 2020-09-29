package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"github.com/yulrizka/grproxy"
)

func main() {
	addr := ":9999"
	target := "127.0.0.1:10000"

	flag.StringVar(&addr, "addr", env("ADDR", addr), "proxy address")
	flag.StringVar(&target, "target", env("TARGET", target), "target grpc server")
	flag.Parse()

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	log.Printf("grproxy %v -> %v", addr, target)

	interceptor := grproxy.NewInterceptor(target)
	if err := interceptor.Serve(context.Background(), l); err != nil {
		log.Fatal(err)
	}
}

func env(s string, def string) string {
	v := os.Getenv(s)
	if v == "" {
		return def
	}
	return v
}
