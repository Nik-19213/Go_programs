package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server struct must embed UnimplementedChatServiceServer to fulfill the interface requirements
type Server struct {
	UnimplementedChatServiceServer
}

// SayHello implements the ChatServiceServer interface
func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received message from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
