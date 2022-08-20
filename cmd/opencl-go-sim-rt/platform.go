package main

// #include "api.h"
import "C"

//export simGetPlatformInfo
func simGetPlatformInfo(_ C.cl_platform_id, paramName C.cl_platform_info,
	paramValueSize C.size_t, paramValue *C.void, paramValueSizeRet *C.size_t) C.cl_int {
	if paramName == C.CL_PLATFORM_ICD_SUFFIX_KHR {
		return stringToCaller("GoSim", paramValueSize, paramValue, paramValueSizeRet)
	} else if paramName == C.CL_PLATFORM_NAME {
		return stringToCaller("github.com/opencl-go/runtime-simulator", paramValueSize, paramValue, paramValueSizeRet)
	}
	return C.CL_INVALID_VALUE
}
