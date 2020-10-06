package chat

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("received message: %s", message.GetBody())

	return &Message{
		Body: fmt.Sprintf("what do you think about %s", message.GetBody()),
	}, nil
}
