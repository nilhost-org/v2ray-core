//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package grpc

import (
	"log"
	"net"

	"fmt"

	"google.golang.org/grpc"
	"v2ray.com/core/app/p2p/wire"
)

const DEFUALT_GRPC_PORT = 55255

func ServiceStart(srv wire.NotifierServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", DEFUALT_GRPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	wire.RegisterNotifierServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
