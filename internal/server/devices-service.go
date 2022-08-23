package server

import (
	"context"

	extSim "github.com/opencl-go/simulator-grpc/sim"
	intSim "github.com/opencl-go/simulator-runtime/internal/sim"
)

// DevicesService provides gRPC access to the platform.
type DevicesService struct {
	extSim.UnimplementedDevicesServer

	platform *intSim.Platform
}

// CreateDevice calls simulator.CreateDevice().
func (service *DevicesService) CreateDevice(_ context.Context, _ *extSim.DeviceCreateRequest) (*extSim.DeviceCreateResponse, error) {
	id := service.platform.CreateDevice()
	return &extSim.DeviceCreateResponse{
		ID: uint64(id),
	}, nil
}
