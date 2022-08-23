package server

import (
	"net"

	extSim "github.com/opencl-go/simulator-grpc/sim"
	intSim "github.com/opencl-go/simulator-runtime/internal/sim"

	"google.golang.org/grpc"
)

// Server represents the entire simulator via gRPC.
type Server struct {
	grpcServer *grpc.Server
	finished   chan error

	address net.Addr
}

// StartServer starts a new server for the given platform.
func StartServer(simPlatform *intSim.Platform) (*Server, error) {
	addr := "127.0.0.1:0"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	// Register listener.Addr()
	grpcServer := grpc.NewServer()
	devicesService := &DevicesService{
		UnimplementedDevicesServer: extSim.UnimplementedDevicesServer{},
		platform:                   simPlatform,
	}
	extSim.RegisterDevicesServer(grpcServer, devicesService)
	server := &Server{
		grpcServer: grpcServer,
		finished:   make(chan error),

		address: listener.Addr(),
	}
	go func() {
		err := server.grpcServer.Serve(listener)
		if err != nil {
			server.finished <- err
		}
		close(server.finished)
	}()
	return server, nil
}

// Stop shuts down the server. Will return when it has done so.
func (server *Server) Stop() {
	server.grpcServer.GracefulStop()
	<-server.finished
}

// Address returns that of the gRPC server.
func (server *Server) Address() net.Addr {
	return server.address
}

// AddressPlatformInfoName returns the info name for the platform to return the address.
func (server *Server) AddressPlatformInfoName() uint32 {
	return extSim.PlatformServerAddressInfo
}
