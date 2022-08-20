package main

// #include "api.h"
import "C"
import "unsafe"

// stringToCaller provides the necessary information of the given string value to the caller pointer values.
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
