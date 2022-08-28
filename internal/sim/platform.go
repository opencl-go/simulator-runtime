package sim

// #include "../api/api.h"
import "C"
import (
	"fmt"
	"strings"
	"sync"
)

// PlatformInfo contains all common platform information fields.
type PlatformInfo struct {
	Name       string
	Vendor     string
	Profile    string
	Version    Version
	Extensions []NameVersion
}

// ExtensionNames returns a list of extension names.
func (info PlatformInfo) ExtensionNames() []string {
	names := make([]string, 0, len(info.Extensions))
	for _, extension := range info.Extensions {
		names = append(names, extension.Name())
	}
	return names
}

// Platform handles all the operations of this runtime simulator.
type Platform struct {
	mutex         sync.Mutex
	allocator     APIObjectAllocator
	info          PlatformInfo
	infoOverrides map[uint32]Information

	object  APIObject
	devices map[ObjectID]*Device
}

// NewPlatform returns a new instance.
func NewPlatform(allocator APIObjectAllocator) *Platform {
	platform := &Platform{
		mutex:     sync.Mutex{},
		allocator: allocator,
		info: PlatformInfo{
			Name:    "runtime-simulator",
			Vendor:  "github.com/opencl-go",
			Profile: "FULL_PROFILE",
			Version: VersionOf(3, 0, 0),
			Extensions: []NameVersion{
				{Version: VersionOf(1, 0, 0), name: "cl_khr_icd"},
			},
		},
		infoOverrides: make(map[uint32]Information),
		devices:       make(map[ObjectID]*Device),
	}
	platform.object = allocator.New(platform)
	return platform
}

// Delete resets the platform and deletes any internal resources.
func (platform *Platform) Delete() {
	platform.mutex.Lock()
	defer platform.mutex.Unlock()

	for _, device := range platform.devices {
		device.Delete()
	}
	platform.devices = make(map[ObjectID]*Device)
	platform.object.Delete()
}

// ID returns the unique identifier of the platform.
func (platform *Platform) ID() ObjectID {
	return platform.object.ID()
}

// Info returns the raw byte value for the queried information.
func (platform *Platform) Info(infoName uint32) Information {
	platform.mutex.Lock()
	defer platform.mutex.Unlock()
	if override, overridden := platform.infoOverrides[infoName]; overridden {
		return override
	}

	switch infoName {
	case C.CL_PLATFORM_PROFILE:
		return InfoString(platform.info.Profile)
	case C.CL_PLATFORM_NAME:
		return InfoString(platform.info.Name)
	case C.CL_PLATFORM_VENDOR:
		return InfoString(platform.info.Vendor)
	case C.CL_PLATFORM_EXTENSIONS:
		return InfoString(strings.Join(platform.info.ExtensionNames(), " "))
	case C.CL_PLATFORM_EXTENSIONS_WITH_VERSION:
		return InfoNameVersion(platform.info.Extensions)
	case C.CL_PLATFORM_HOST_TIMER_RESOLUTION:
		return InfoUint64(0)
	case C.CL_PLATFORM_VERSION:
		return InfoString(fmt.Sprintf("OpenCL %d.%d ", platform.info.Version.Major(), platform.info.Version.Minor()))
	case C.CL_PLATFORM_NUMERIC_VERSION:
		return InfoVersion(platform.info.Version)
	}
	return nil
}

// SetInfo sets and overrides any information the platform provides.
func (platform *Platform) SetInfo(infoName uint32, value Information) {
	platform.mutex.Lock()
	defer platform.mutex.Unlock()
	platform.infoOverrides[infoName] = value
}

// ResetInfo clears any information override.
func (platform *Platform) ResetInfo(infoName uint32) {
	platform.mutex.Lock()
	defer platform.mutex.Unlock()
	delete(platform.infoOverrides, infoName)
}

// CreateDevice adds a new device to the platform; The returned value is the device ID.
// With the returned values, an OpenCL client can determine which device of all returned the one created is.
func (platform *Platform) CreateDevice() ObjectID {
	platform.mutex.Lock()
	defer platform.mutex.Unlock()

	device := NewDevice(platform.ID(), platform.allocator)
	platform.devices[device.ID()] = device
	return device.ID()
}

// Devices returns all available devices. The order of the returned list may vary.
func (platform *Platform) Devices() []*Device {
	platform.mutex.Lock()
	defer platform.mutex.Unlock()

	devices := make([]*Device, 0, len(platform.devices))
	for _, device := range platform.devices {
		devices = append(devices, device)
	}
	return devices
}

// DeleteDevice removes the identified device from the platform.
func (platform *Platform) DeleteDevice(id ObjectID) {
	platform.mutex.Lock()
	defer platform.mutex.Unlock()

	device := platform.devices[id]
	delete(platform.devices, id)
	device.Delete()
}
