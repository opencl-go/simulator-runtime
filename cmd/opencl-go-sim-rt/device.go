package main

// #include "api.h"
import "C"
import (
	"unsafe"

	"github.com/opencl-go/simulator-runtime/internal/sim"
)

//export goGetDeviceIDs
func goGetDeviceIDs(platformHandle C.uint64_t, deviceType C.cl_device_type,
	deviceCount C.cl_uint, devicesPtr *C.cl_device_id, deviceCountRet *C.cl_uint) C.cl_int {
	if ((deviceCount > 0) && (devicesPtr == nil)) || ((devicesPtr == nil) && (deviceCountRet == nil)) {
		return C.CL_INVALID_VALUE
	}
	platform := DispatchObjectFrom(uint64(platformHandle)).Ref().(*sim.Platform)
	deviceObjects := platform.Devices()
	knownIDs := make([]uintptr, len(deviceObjects))
	for i, device := range deviceObjects {
		knownIDs[i] = uintptr(device.ID())
	}
	if deviceCountRet != nil {
		*deviceCountRet = C.cl_uint(len(knownIDs))
	}
	if deviceCount > 0 {
		deviceIDs := unsafe.Slice((*uintptr)(unsafe.Pointer(devicesPtr)), int(deviceCount))
		for i := uint32(0); i < uint32(deviceCount); i++ {
			deviceIDs[i] = 0
		}
		copy(deviceIDs, knownIDs)
	}
	return C.CL_SUCCESS
}

//export goGetDeviceInfo
func goGetDeviceInfo(deviceHandle C.uint64_t,
	paramName C.cl_device_info, paramValueSize C.size_t, paramValue *C.void, paramValueSizeRet *C.size_t) C.cl_int {
	device := DispatchObjectFrom(uint64(deviceHandle)).Ref().(*sim.Device)
	if info := device.Info(uint32(paramName)); info != nil {
		return dataToCaller(info.Value(), paramValueSize, paramValue, paramValueSizeRet)
	}
	return C.CL_INVALID_VALUE
}
