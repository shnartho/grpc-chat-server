// chat_server/chat_server.go
package main

import (
	"log"
	"net"

	pb "chat-grpc" // Import generated code

	"google.golang.org/grpc"
)

type chatServer struct {
	pb.UnimplementedChatServer
}

func (s *chatServer) Join(stream pb.Chat_JoinServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("[%s]: %s\n", msg.GetUser(), msg.GetMessage())

		// Send the received message back to the client with a newline character
		err = stream.Send(&pb.ChatMessage{
			User:    msg.GetUser(),
			Message: msg.GetMessage() + "\n",
		})
		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServer(s, &chatServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
