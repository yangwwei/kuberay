package manager

import (
	"fmt"

	"github.com/ray-project/kuberay/apiserver/pkg/apis/generated/go/workloadmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AppleBatchClient struct {
	workloadmanager.WorkloadManagerClient
	connection *grpc.ClientConn
}

func NewAppleBatchClient(address string) (*AppleBatchClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	fmt.Printf("workload client connection target: %s\n", conn.Target())
	workloadClient := workloadmanager.NewWorkloadManagerClient(conn)
	return &AppleBatchClient{
		WorkloadManagerClient: workloadClient,
		connection:            conn,
	}, nil
}

func (w *AppleBatchClient) Close() {
	fmt.Println("closing workload client")
	err := w.connection.Close()
	if err != nil {
		_ = fmt.Errorf("failed to close the connection for the workload client, error=%v", err)
	}
}
