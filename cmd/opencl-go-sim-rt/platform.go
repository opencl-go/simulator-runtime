package main

// #include "api.h"
import "C"
import (
	"sync"
	"unsafe"

	"github.com/opencl-go/simulator-runtime/internal/server"
	"github.com/opencl-go/simulator-runtime/internal/sim"
)

var (
	startup           sync.Once
	simulatorPlatform *sim.Platform
	rpcServer         *server.Server
)

type objectAllocatorFunc func(ref any) *DispatchObject

func (f objectAllocatorFunc) New(ref any) sim.APIObject {
	return f(ref)
}

//export goGetPlatformIDs
func goGetPlatformIDs(platformCount C.cl_uint, platformsPtr *C.cl_platform_id, platformCountRet *C.cl_uint) C.cl_int {
	if ((platformCount > 0) && (platformsPtr == nil)) || ((platformsPtr == nil) && (platformCountRet == nil)) {
		return C.CL_INVALID_VALUE
	}
	if platformCountRet != nil {
		*platformCountRet = 1
	}
	if platformCount > 0 {
		platforms := unsafe.Slice((*uintptr)(unsafe.Pointer(platformsPtr)), uint32(platformCount))
		for i := uint32(0); i < uint32(platformCount); i++ {
			platforms[i] = 0
		}
		startup.Do(func() {
			simulatorPlatform = sim.NewPlatform(objectAllocatorFunc(NewDispatchObject))
			var err error
			rpcServer, err = server.StartServer(simulatorPlatform)
			var serverAddress string
			if err == nil {
				serverAddress = rpcServer.Address().String()
			}
			simulatorPlatform.SetInfo(rpcServer.AddressPlatformInfoName(), sim.InfoString(serverAddress))
		})
		platforms[0] = uintptr(simulatorPlatform.ID())
	}
	return C.CL_SUCCESS
}

//export goGetPlatformInfo
func goGetPlatformInfo(platformHandle C.uint64_t, paramName C.cl_platform_info,
	paramValueSize C.size_t, paramValue *C.void, paramValueSizeRet *C.size_t) C.cl_int {
	platform := DispatchObjectFrom(uint64(platformHandle)).Ref().(*sim.Platform)
	if paramName == C.CL_PLATFORM_ICD_SUFFIX_KHR {
		return dataToCaller(sim.InfoString("GoSim").Value(), paramValueSize, paramValue, paramValueSizeRet)
	} else if info := platform.Info(uint32(paramName)); info != nil {
		return dataToCaller(info.Value(), paramValueSize, paramValue, paramValueSizeRet)
	}
	return C.CL_INVALID_VALUE
}

//export goGetExtensionFunctionAddressForPlatform
func goGetExtensionFunctionAddressForPlatform(platformHandle C.uint64_t, rawFuncName *C.char) C.uintptr_t {
	funcName := C.GoString(rawFuncName)
	platform := DispatchObjectFrom(uint64(platformHandle)).Ref().(*sim.Platform)
	return C.uintptr_t(platform.ExtensionFunctionAddress(funcName))
}
