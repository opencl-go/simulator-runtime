package sim

// #include "../api/api.h"
import "C"
import (
	"fmt"
	"strings"
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
	info          PlatformInfo
	infoOverrides map[uint32]Information
}

// NewPlatform returns a new instance.
func NewPlatform() *Platform {
	platform := Platform{
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
	}
	return &platform
}

// Info returns the raw byte value for the queried information.
func (p *Platform) Info(infoName uint32) Information {
	if override, overridden := p.infoOverrides[infoName]; overridden {
		return override
	}

	switch infoName {
	case C.CL_PLATFORM_PROFILE:
		return InfoString(p.info.Profile)
	case C.CL_PLATFORM_NAME:
		return InfoString(p.info.Name)
	case C.CL_PLATFORM_VENDOR:
		return InfoString(p.info.Vendor)
	case C.CL_PLATFORM_EXTENSIONS:
		return InfoString(strings.Join(p.info.ExtensionNames(), " "))
	case C.CL_PLATFORM_EXTENSIONS_WITH_VERSION:
		return InfoNameVersion(p.info.Extensions)
	case C.CL_PLATFORM_HOST_TIMER_RESOLUTION:
		return InfoUint64(0)
	case C.CL_PLATFORM_VERSION:
		return InfoString(fmt.Sprintf("OpenCL %d.%d ", p.info.Version.Major(), p.info.Version.Minor()))
	case C.CL_PLATFORM_NUMERIC_VERSION:
		return InfoVersion(p.info.Version)
	}
	return nil
}

// SetInfo sets and overrides any information the platform provides.
func (p *Platform) SetInfo(infoName uint32, value Information) {
	p.infoOverrides[infoName] = value
}

// ResetInfo clears any information override.
func (p *Platform) ResetInfo(infoName uint32) {
	delete(p.infoOverrides, infoName)
}
