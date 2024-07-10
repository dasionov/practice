package main

import (
	"log"
	"net"
	"practice/pkg/server"

	"google.golang.org/grpc"
	pb "practice/pkg/proto"
)

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	pb.RegisterServerServer(s, &server.Server{})

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
