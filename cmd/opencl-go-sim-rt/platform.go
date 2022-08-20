package main

// #include "api.h"
import "C"
import (
	"runtime/cgo"
	"sync"

	"github.com/opencl-go/simulator-runtime/internal/sim"
)

var (
	startup                 sync.Once
	simulatorPlatform       *sim.Platform
	simulatorPlatformHandle cgo.Handle
)

//export simGetPlatform
func simGetPlatform() C.uintptr_t {
	startup.Do(func() {
		simulatorPlatform = sim.NewPlatform()
		simulatorPlatformHandle = cgo.NewHandle(simulatorPlatform)
	})
	return C.uintptr_t(uintptr(simulatorPlatformHandle))
}

//export simGetPlatformInfo
func simGetPlatformInfo(platformHandle C.uintptr_t, paramName C.cl_platform_info,
	paramValueSize C.size_t, paramValue *C.void, paramValueSizeRet *C.size_t) C.cl_int {
	platform := cgo.Handle(platformHandle).Value().(*sim.Platform)
	if paramName == C.CL_PLATFORM_ICD_SUFFIX_KHR {
		return dataToCaller(sim.InfoString("GoSim").Value(), paramValueSize, paramValue, paramValueSizeRet)
	} else if value, known := platform.Info(uint32(paramName)); known {
		return dataToCaller(value, paramValueSize, paramValue, paramValueSizeRet)
	}
	return C.CL_INVALID_VALUE
}
