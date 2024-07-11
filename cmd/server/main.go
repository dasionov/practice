package main

import (
	"fmt"
	"log"
	"net"
	"practice/pkg/server"
	"practice/pkg/virt-launcher/virtwrap"
	"time"

	"google.golang.org/grpc"
	pb "practice/pkg/proto"
)

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	domainConn := createLibvirtConnection(*runWithNonRoot)
	defer domainConn.Close()

	domainManager, err := virtwrap.NewLibvirtDomainManager(domainConn, *virtShareDir, *ephemeralDiskDir, &agentStore, *ovmfPath, ephemeralDiskCreator, metadataCache)
	if err != nil {
		log.Fatalf("Failed to create domain manager: %v", err)
	}

	// Initialize the server with the domain manager
	s := &server.Server{
		domainManager:  domainManager,
		allowEmulation: false, // Set this based on your requirements
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterServerServer(grpcServer, s)

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func createLibvirtConnection(runWithNonRoot bool) virtcli.Connection {
	libvirtUri := "qemu:///system"
	user := ""
	if runWithNonRoot {
		user = putil.NonRootUserString
		libvirtUri = "qemu+unix:///session?socket=/var/run/libvirt/virtqemud-sock"
	}

	domainConn, err := virtcli.NewConnection(libvirtUri, user, "", 10*time.Second)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to virtqemud: %v", err))
	}

	return domainConn
}
