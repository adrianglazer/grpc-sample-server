package main

import (
	"grpc-sample-server/message"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	m := message.Server{}
	grpcServer := grpc.NewServer()

	message.RegisterMessageServiceServer(grpcServer, &m)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
