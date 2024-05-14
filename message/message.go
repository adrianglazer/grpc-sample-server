package message

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedMessageServiceServer
}

func (s *Server) Hello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received message from server: %s\n", in.Body)

	return &Message{Body: "Hello from the server"}, nil
}
