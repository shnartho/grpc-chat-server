// chat_client/chat_client.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "chat-grpc" // Import generated code

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatClient(conn)

	ctx := context.Background()

	stream, err := client.Join(ctx)
	if err != nil {
		log.Fatalf("could not join: %v", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("error receiving message: %v", err)
			}
			log.Printf("[%s]: %s", msg.GetUser(), msg.GetMessage())
		}
	}()

	for {
		var message string
		if _, err := fmt.Scanln(&message); err != nil {
			log.Fatalf("error reading message: %v", err)
		}

		if message == "/exit" {
			os.Exit(0)
		}

		err := stream.Send(&pb.ChatMessage{
			User:    "User",
			Message: message,
		})
		if err != nil {
			log.Fatalf("error sending message: %v", err)
		}
		time.Sleep(time.Second)
	}
}
