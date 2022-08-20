package main

// #include "api.h"
import "C"
import "unsafe"

func main() {}

func stringToCaller(value string, paramValueSize C.size_t, paramValue *C.void, paramValueSizeRet *C.size_t) C.cl_int {
	valueBytes := []byte(value)
	requiredSize := len(valueBytes) + 1
	if paramValueSizeRet != nil {
		*paramValueSizeRet = (C.size_t)(requiredSize)
	}
	if paramValue == nil {
		return C.CL_SUCCESS
	}
	if paramValueSize < C.size_t(requiredSize) {
		return C.CL_INVALID_VALUE
	}
	paramValueBytes := unsafe.Slice((*byte)(unsafe.Pointer(paramValue)), uintptr(paramValueSize))
	copy(paramValueBytes, valueBytes)
	paramValueBytes[requiredSize-1] = 0x00
	return C.CL_SUCCESS
}

//export simGetPlatformInfo
func simGetPlatformInfo(platform C.cl_platform_id, paramName C.cl_platform_info,
	paramValueSize C.size_t, paramValue *C.void, paramValueSizeRet *C.size_t) C.cl_int {
	if paramName == C.CL_PLATFORM_ICD_SUFFIX_KHR {
		return stringToCaller("GoSim", paramValueSize, paramValue, paramValueSizeRet)
	} else if paramName == C.CL_PLATFORM_NAME {
		return stringToCaller("github.com/opencl-go/runtime-simulator", paramValueSize, paramValue, paramValueSizeRet)
	}
	return C.CL_INVALID_VALUE
}
