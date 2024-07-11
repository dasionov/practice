package main

import (
	"context"
	"log"
	"practice/pkg/client"
	pb "practice/pkg/proto"
	"time"
)

func main() {
	// Prompt user for the file path
	filePath := client.ReadFilePath()

	// Set a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Create a gRPC connection
	conn, err := client.CreateGRPCConnection("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	grpcClient := pb.NewServerClient(conn)

	// Process the file
	err = client.ProcessFile(filePath, grpcClient, ctx)
	if err != nil {
		log.Printf("Error processing file %s: %v", filePath, err)
	}
}
