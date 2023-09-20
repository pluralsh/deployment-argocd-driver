package driver

import (
	"context"

	deploymentspec "github.com/pluralsh/deployment-api/spec"
	"github.com/pluralsh/deployment-argocd-driver/pkg/argocd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewDriver(provisioner string) (*IdentityServer, *ProvisionerServer) {
	return &IdentityServer{
			provisioner: provisioner,
		}, &ProvisionerServer{
			provisioner: provisioner,
		}
}

type ProvisionerServer struct {
	provisioner string
	argocd      *argocd.Argocd
}

func (ps *ProvisionerServer) DriverGetDeploymentStatus(ctx context.Context, request *deploymentspec.DriverGetDeploymentStatusRequest) (*deploymentspec.DriverGetDeploymentStatusResponse, error) {
	return &deploymentspec.DriverGetDeploymentStatusResponse{
		DeploymentId: request.DeploymentId,
		DeploymentStatus: &deploymentspec.DeploymentStatusEnum{
			Type: &deploymentspec.DeploymentStatusEnum_Ready{
				Ready: true,
			},
		},
		Message: "",
	}, nil
}

func (ps *ProvisionerServer) DriverCreateDeployment(_ context.Context, req *deploymentspec.DriverCreateDeploymentRequest) (*deploymentspec.DriverCreateDeploymentResponse, error) {
	deploymentName := req.GetName()

	dbID := deploymentName

	return &deploymentspec.DriverCreateDeploymentResponse{
		DeploymentId: dbID,
	}, nil
}

func (ps *ProvisionerServer) DriverDeleteDeployment(_ context.Context, req *deploymentspec.DriverDeleteDeploymentRequest) (*deploymentspec.DriverDeleteDeploymentResponse, error) {

	return &deploymentspec.DriverDeleteDeploymentResponse{}, status.Error(codes.NotFound, "Deployment not found")
}

type IdentityServer struct {
	provisioner string
}

func (id *IdentityServer) DriverGetInfo(context.Context, *deploymentspec.DriverGetInfoRequest) (*deploymentspec.DriverGetInfoResponse, error) {
	if id.provisioner == "" {
		return nil, status.Error(codes.InvalidArgument, "ProvisionerName is empty")
	}

	return &deploymentspec.DriverGetInfoResponse{
		Name: id.provisioner,
	}, nil
}
