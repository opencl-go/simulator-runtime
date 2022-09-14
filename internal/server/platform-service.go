package server

import (
	"context"

	extSim "github.com/opencl-go/simulator-grpc/sim"
	intSim "github.com/opencl-go/simulator-runtime/internal/sim"
)

// PlatformService provides gRPC access to the platform.
type PlatformService struct {
	extSim.UnimplementedPlatformServer

	platform *intSim.Platform
}

// PrepareExtensionFunction calls Platform.PrepareExtensionFunction().
func (service *PlatformService) PrepareExtensionFunction(
	ctx context.Context,
	request *extSim.PlatformPrepareExtensionFunctionRequest) (
	*extSim.PlatformPrepareExtensionFunctionResponse,
	error) {
	name, addr := service.platform.PrepareExtensionFunction()
	response := &extSim.PlatformPrepareExtensionFunctionResponse{
		Name:    name,
		Address: uint64(addr),
	}
	return response, nil
}

// ReleaseExtensionFunction calls Platform.ReleaseExtensionFunction().
func (service *PlatformService) ReleaseExtensionFunction(
	ctx context.Context,
	request *extSim.PlatformReleaseExtensionFunctionRequest) (
	*extSim.PlatformReleaseExtensionFunctionResponse,
	error) {
	service.platform.ReleaseExtensionFunction(request.Name)
	return &extSim.PlatformReleaseExtensionFunctionResponse{}, nil
}
