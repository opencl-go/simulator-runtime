package main

// #include "api.h"
import "C"
import "github.com/opencl-go/simulator-runtime/internal/sim"

//export simGetPlatformInfo
func simGetPlatformInfo(_ C.cl_platform_id, paramName C.cl_platform_info,
	paramValueSize C.size_t, paramValue *C.void, paramValueSizeRet *C.size_t) C.cl_int {
	if paramName == C.CL_PLATFORM_ICD_SUFFIX_KHR {
		return dataToCaller(sim.InfoString("GoSim").Value(), paramValueSize, paramValue, paramValueSizeRet)
	} else if value, known := sim.PlatformInfo(uint32(paramName)); known {
		return dataToCaller(value, paramValueSize, paramValue, paramValueSizeRet)
	}
	return C.CL_INVALID_VALUE
}
