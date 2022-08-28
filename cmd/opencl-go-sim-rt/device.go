package main

// #include "api.h"
import "C"
import "unsafe"

//export goGetDeviceIDs
func goGetDeviceIDs(platformHandle C.uint64_t, deviceType C.cl_device_type,
	deviceCount C.cl_uint, devices *C.cl_device_id, deviceCountRet *C.cl_uint) C.cl_int {
	return C.CL_OUT_OF_RESOURCES
}

//export goGetDeviceInfo
func goGetDeviceInfo(deviceHandle C.uint64_t,
	paramName C.cl_device_info, paramValueSize C.size_t, paramValue unsafe.Pointer, paramValueSizeRet *C.size_t) C.cl_int {
	return C.CL_OUT_OF_RESOURCES
}
