package main

import (
	"github.com/rajatparida86/grpcChatService/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("listening on port 9000...")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000 - %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start server - %v", err)
	}
}
