package main

import (
	"log"
	"net"

	v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"google.golang.org/grpc"

	"github.com/dio/alsserver/pkg/als"
)

func main() {
	grpcServer := grpc.NewServer()
	v2.RegisterAccessLogServiceServer(grpcServer, als.New())

	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening on tcp://localhost:9001")
	grpcServer.Serve(l)
}
