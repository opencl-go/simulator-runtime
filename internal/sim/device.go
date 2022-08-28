package sim

import "C"
import (
	"sync"
)

// Device simulates the behaviour of one OpenCL device.
type Device struct {
	mutex      sync.Mutex
	platformID ObjectID
	object     APIObject

	infoOverrides map[uint32]Information
}

// NewDevice creates a new instance.
func NewDevice(platformID ObjectID, allocator APIObjectAllocator) *Device {
	device := &Device{
		mutex:      sync.Mutex{},
		platformID: platformID,
	}
	device.object = allocator.New(device)
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
	// TODO: any fixed info
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
