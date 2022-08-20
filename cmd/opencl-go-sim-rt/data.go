package main

// #include "api.h"
import "C"
import "unsafe"

// dataToCaller provides the necessary information of the given byte slice value to the caller pointer values.
func dataToCaller(value []byte, paramValueSize C.size_t, paramValue *C.void, paramValueSizeRet *C.size_t) C.cl_int {
	requiredSize := len(value)
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
	copy(paramValueBytes, value)
	return C.CL_SUCCESS
}
