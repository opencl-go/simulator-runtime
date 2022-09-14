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

// Create calls Platform.CreateDevice().
func (service *DevicesService) Create(_ context.Context, _ *extSim.DeviceCreateRequest) (*extSim.DeviceCreateResponse, error) {
	id := service.platform.CreateDevice()
	return &extSim.DeviceCreateResponse{
		ID: uint64(id),
	}, nil
}

// Delete calls Platform.DeleteDevice().
func (service *DevicesService) Delete(_ context.Context, request *extSim.DeviceDeleteRequest) (*extSim.DeviceDeleteResponse, error) {
	service.platform.DeleteDevice(intSim.ObjectID(request.ID))
	return &extSim.DeviceDeleteResponse{}, nil
}
