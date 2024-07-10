package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "kubevirt.io/api/core/v1"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "practice/pkg/proto"
)

func ReadFilePath() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the path of the YAML file: ")
	filePath, _ := reader.ReadString('\n')
	return filePath[:len(filePath)-1] // remove newline character
}

func ReadYAMLFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func UnmarshalYAML(fileContent []byte, vmi *v1.VirtualMachineInstance) error {
	return yaml.Unmarshal(fileContent, vmi)
}

func CreateGRPCConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func SyncVirtualMachine(client pb.ServerClient, ctx context.Context, vmiJson []byte) (*pb.Response, error) {
	req := &pb.VMIRequest{
		Vmi: &pb.VMI{VmiJson: vmiJson},
	}
	return client.SyncVirtualMachine(ctx, req)
}

func ProcessFile(filePath string, client pb.ServerClient, ctx context.Context) error {
	fileContent, err := ReadYAMLFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var vmi v1.VirtualMachineInstance
	err = UnmarshalYAML(fileContent, &vmi)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	vmiJson, err := json.Marshal(vmi)
	if err != nil {
		return fmt.Errorf("failed to marshal VMI to JSON: %w", err)
	}

	res, err := SyncVirtualMachine(client, ctx, vmiJson)
	if err != nil {
		return fmt.Errorf("failed to sync VMI: %w", err)
	}

	log.Printf("Response from server: %s", res.Message)
	return nil
}
