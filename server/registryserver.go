package server

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/open-beagle/awecloud-bmq-sdk/pkg"
	"github.com/open-beagle/awecloud-bmq-server/pkg/conf"
	"github.com/open-beagle/awecloud-bmq-server/pkg/data"
)

func NewRegistryServer() {
	lis, err := net.Listen("tcp", conf.GRPC.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pkg.RegisterRegistryServer(s, &RegistryServer{})
	log.Printf("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// RegistryServer is used to implement sdk.RegistryServer.
type RegistryServer struct {
	pkg.UnimplementedRegistryServer
}

// Login implements sdk.RegistryServer.
func (s *RegistryServer) Login(ctx context.Context, in *pkg.LoginRequest) (*pkg.LoginResponse, error) {
	worker := conf.Server.GetWorker(in.ID)
	if worker == nil {
		return nil, errors.New("error : Worker do not exist.")
	}
	if worker.Secret == in.Secret {
		return &pkg.LoginResponse{
			Path:   conf.Message.Prefix,
			Secret: conf.Message.Token,
		}, nil
	}
	return nil, errors.New("error : Worker Secret mismatch")
}

// Listen implements sdk.RegistryServer.
func (s *RegistryServer) Listen(in *pkg.ListenRequest, stream pkg.Registry_ListenServer) error {
	worker := &data.OnlineWoker{
		ID:     in.ID,
		Kind:   in.Kind,
		OS:     in.OS,
		Arch:   in.Arch,
		Kernel: in.Kernel,
	}
	data.Server.SetWorker(worker)

	// Start a ticker that executes each 5 seconds
	timer := time.NewTicker(5 * time.Second)
	result := &pkg.ListenResponse{}
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			// Grab stats and output
			result = &pkg.ListenResponse{
				Action: pkg.Action_GetServices,
			}
			err := stream.Send(result)
			if err != nil {
				return err
			}
		case <-worker.Channel:
			result = &pkg.ListenResponse{
				Action: pkg.Action_GetServices,
			}
			err := stream.Send(result)
			if err != nil {
				return err
			}
		}
	}
}

// GetServices implements sdk.RegistryServer.
func (s *RegistryServer) GetServices(ctx context.Context, in *pkg.GetServicesRequest) (*pkg.GetServicesResponse, error) {
	result := &pkg.GetServicesResponse{}
	worker := conf.Server.GetWorker(in.ID)
	if worker != nil {
		if len(worker.Agents) > 0 {
			result.AgentServices = make([]*pkg.ServiceConfig, len(worker.Agents))
			for k, agent := range worker.Agents {
				if len(agent.Services) > 0 {
					result.AgentServices[k] = &pkg.ServiceConfig{
						Server_Address: conf.Host,
						Server_Port:    strconv.Itoa(conf.Port),
						User:           agent.User,
						Token:          conf.Message.Token,
						Services:       make(map[string]*pkg.SecretService),
					}
					for _, service := range agent.Services {
						result.AgentServices[k].Services[service.Name] = &pkg.SecretService{
							Type:   service.Type,
							Name:   service.Name,
							Secret: service.SK,
							Host:   service.Local_IP,
							Port:   service.Local_Port,
						}
					}
				}
			}
		}
		if len(worker.Visitors) > 0 {
			result.VisitorServices = make([]*pkg.ServiceConfig, len(worker.Visitors))
			for k, visitor := range worker.Visitors {
				if len(visitor.Services) > 0 {
					result.AgentServices[k] = &pkg.ServiceConfig{
						Server_Address: conf.Host,
						Server_Port:    strconv.Itoa(conf.Port),
						User:           visitor.User,
						Token:          conf.Message.Token,
						Services:       make(map[string]*pkg.SecretService),
					}
					for _, service := range visitor.Services {
						result.AgentServices[k].Services[service.Name] = &pkg.SecretService{
							Type:   service.Type,
							Name:   service.Name,
							Secret: service.SK,
							Host:   service.Bind_Addr,
							Port:   service.Bind_Port,
						}
					}
				}
			}
		}
	}
	return result, nil
}
