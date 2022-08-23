package sim

import "sync"

// Device simulates the behaviour of one OpenCL device.
type Device struct {
	mutex      sync.Mutex
	platformID ObjectID
	object     APIObject
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
