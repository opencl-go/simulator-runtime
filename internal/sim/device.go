package sim

// #include "../api/api.h"
import "C"
import (
	"fmt"
	"sync"
)

// DeviceInfo contains all common platform information fields.
type DeviceInfo struct {
	PlatformID ObjectID
	Name       string
	Vendor     string
}

// Device simulates the behaviour of one OpenCL device.
type Device struct {
	mutex  sync.Mutex
	object APIObject
	info   DeviceInfo

	infoOverrides map[uint32]Information
}

// NewDevice creates a new instance.
func NewDevice(platformID ObjectID, allocator APIObjectAllocator) *Device {
	device := &Device{
		mutex: sync.Mutex{},
		info: DeviceInfo{
			PlatformID: platformID,
			Vendor:     "github.com/opencl-go",
		},
	}
	device.object = allocator.New(device)
	device.info.Name = fmt.Sprintf("Dynamic Device %v", device.object.ID())
	return device
}

// Delete cleans up the device.
func (device *Device) Delete() {
	if device == nil {
		return
	}
	device.mutex.Lock()
	defer device.mutex.Unlock()
	device.object.Delete()
}

// ID returns the API identifier of the device.
func (device *Device) ID() ObjectID {
	device.mutex.Lock()
	defer device.mutex.Unlock()
	return device.object.ID()
}

// Info returns the raw byte value for the queried information.
func (device *Device) Info(infoName uint32) Information {
	device.mutex.Lock()
	defer device.mutex.Unlock()
	if override, overridden := device.infoOverrides[infoName]; overridden {
		return override
	}
	switch infoName {
	case C.CL_DEVICE_PLATFORM:
		return InfoUintptr(device.info.PlatformID)
	case C.CL_DEVICE_VENDOR:
		return InfoString(device.info.Vendor)
	case C.CL_DEVICE_NAME:
		return InfoString(device.info.Name)
	}
	return nil
}

// SetInfo sets and overrides any information the device provides.
func (device *Device) SetInfo(infoName uint32, value Information) {
	device.mutex.Lock()
	defer device.mutex.Unlock()
	device.infoOverrides[infoName] = value
}

// ResetInfo clears any information override.
func (device *Device) ResetInfo(infoName uint32) {
	device.mutex.Lock()
	defer device.mutex.Unlock()
	delete(device.infoOverrides, infoName)
}
