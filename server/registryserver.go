package server

import (
	"context"

	"github.com/open-beagle/awecloud-bmq-sdk/pkg"
)

// RegistryServer is used to implement sdk.RegistryServer.
type RegistryServer struct {
	pkg.UnimplementedRegistryServer
	ListenChan chan *pkg.ListenResponse
}

// Login implements sdk.RegistryServer.
func (s *RegistryServer) Login(ctx context.Context, in *pkg.LoginRequest) (*pkg.LoginResponse, error) {
	return nil, nil
}

// Listen implements sdk.RegistryServer.
func (s *RegistryServer) Listen(in *pkg.ListenRequest, stream pkg.Registry_ListenServer) error {
	for {
		res := <-s.ListenChan
		stream.Send(res)
	}
	return nil
}

// GetServices implements sdk.RegistryServer.
func (s *RegistryServer) GetServices(ctx context.Context, in *pkg.GetServicesRequest) (*pkg.GetServicesResponse, error) {
	return nil, nil
}

type registryWoker struct {
	kind    string
	os      string
	arch    string
	kernel  string
	variant string
	labels  map[string]string
	channel chan *pkg.ListenResponse
}
