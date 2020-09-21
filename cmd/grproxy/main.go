package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/yulrizka/grproxy"
)

func main() {
	var addr, target string
	flag.StringVar(&addr, "addr", "127.0.0.1:9999", "proxy address")
	flag.StringVar(&target, "target", "127.0.0.1:10000", "target grpc server")
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
