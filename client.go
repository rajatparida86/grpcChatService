package main

import (
	"github.com/rajatparida86/grpcChatService/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer conn.Close()

	client := chat.NewChatServiceClient(conn)

	resp, err := client.SayHello(context.Background(), &chat.Message{
		Body: "I dont think",
	})
	if err != nil {
		log.Printf("failed to get response from server: %s", err)
	} else {
		log.Printf("received message: %s", resp.GetBody())
	}
}
