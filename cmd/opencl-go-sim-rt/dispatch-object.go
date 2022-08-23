package main

// #include "icd.h"
import "C"
import (
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/opencl-go/simulator-runtime/internal/sim"
)

var (
	dispatchObjectCounter uint64
	dispatchObjects       = sync.Map{} // map[uint64]*DispatchObject
)

// DispatchObject is an ICD-compatible object that has fixed memory allocated storing the dispatch table.
type DispatchObject struct {
	ptr *C.goDispatchObject
	ref any
}

// NewDispatchObject returns a new instance.
func NewDispatchObject(ref any) *DispatchObject {
	handle := atomic.AddUint64(&dispatchObjectCounter, 1)
	obj := &DispatchObject{ptr: C.newDispatchObject(C.uint64_t(handle)), ref: ref}
	dispatchObjects.Store(handle, obj)
	return obj
}

// DispatchObjectFrom resolves a dispatch object from given handle.
func DispatchObjectFrom(handle uint64) *DispatchObject {
	obj, _ := dispatchObjects.Load(handle)
	return obj.(*DispatchObject)
}

// Delete deallocates the memory and clears the handle association.
func (obj *DispatchObject) Delete() {
	if (obj != nil) && (obj.ptr != nil) {
		dispatchObjects.Delete(obj.ptr.handle)
		C.free(unsafe.Pointer(obj.ptr))
		obj.ptr = nil
	}
}

// ID returns the ID equivalent of the object. It is based on the address of the allocated memory.
func (obj *DispatchObject) ID() sim.ObjectID {
	if obj == nil {
		return 0
	}
	return sim.ObjectID(unsafe.Pointer(obj.ptr))
}

// Handle returns the internal reference to resolve the object.
func (obj *DispatchObject) Handle() uint64 {
	if (obj == nil) || (obj.ptr == nil) {
		return 0
	}
	return uint64(obj.ptr.handle)
}

// Ref returns the referred-to object.
func (obj *DispatchObject) Ref() any {
	if obj == nil {
		return nil
	}
	return obj.ref
}
