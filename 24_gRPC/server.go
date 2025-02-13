package main

//go get google.golang.org/protobuf@latest
//go get google.golang.org/grpc
// protoc --go_out=chat --go-grpc_out=chat chat.proto

import (
	"fmt"
	"log"
	"net"

	"grpctutorial/chat/chat"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
