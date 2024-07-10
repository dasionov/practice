package server

import (
	"context"
	"encoding/json"
	v1 "kubevirt.io/api/core/v1"
	pb "practice/pkg/proto"
)

// Server struct implements the gRPC server interface
type Server struct {
	pb.UnimplementedServerServer
}

// SyncVirtualMachine handles the incoming VMIRequest
func (s *Server) SyncVirtualMachine(ctx context.Context, req *pb.VMIRequest) (*pb.Response, error) {
	vmi := getVMIFromRequest(req.Vmi)

	// Convert to XML

	return &pb.Response{Message: "VMI processed successfully"}, nil
}

func getVMIFromRequest(request *pb.VMI) *v1.VirtualMachineInstance {

	var vmi v1.VirtualMachineInstance
	json.Unmarshal(request.VmiJson, &vmi)
	return &vmi
}
