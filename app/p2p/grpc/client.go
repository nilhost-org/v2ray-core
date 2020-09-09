// Package main implements a client for Greeter service.
package grpc

import (
	"context"
	"log"
	"time"

	"fmt"

	"google.golang.org/grpc"
	"v2ray.com/core/app/p2p/wire"
)

func ClientStart(seed *wire.HelloSeedList) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", DEFUALT_GRPC_PORT), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := wire.NewNotifierClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err = c.SayHello(ctx, seed)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	cancel()
}
