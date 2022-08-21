package sim

// #include "../api/api.h"
import "C"
import (
	"encoding/binary"
	"unsafe"
)

var hostByteOrder binary.ByteOrder

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)
	switch buf {
	case [2]byte{0xCD, 0xAB}:
		hostByteOrder = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		hostByteOrder = binary.BigEndian
	default:
		panic("could not determine native endianness")
	}
}

// Platform handles all the operations of this runtime simulator.
type Platform struct {
}

// Information is a generic provider of data.
type Information interface {
	Value() []byte
}

// InfoString specializes the Information type by representing a char[] information in the API.
type InfoString string

// Value returns the string value as UTF-8 byte array, with a terminating NUL character.
func (value InfoString) Value() []byte {
	return append([]byte(value), 0x00)
}

// InfoUint64 specializes the Information type by representing a 64 bit value.
type InfoUint64 uint64

// Value returns the number serialized in host byte order.
func (value InfoUint64) Value() []byte {
	bytes := make([]byte, 8)
	hostByteOrder.PutUint64(bytes, uint64(value))
	return bytes
}

// InfoZero specializes the Information type as a marker of empty information.
type InfoZero struct{}

// Value returns nil (an empty slice).
func (InfoZero) Value() []byte {
	return nil
}

// InfoVersion specializes the Information type as a Version.
type InfoVersion Version

// Value returns the serialized form of the version in host byte order.
func (value InfoVersion) Value() []byte {
	bytes := make([]byte, 4)
	hostByteOrder.PutUint32(bytes, uint32(value))
	return bytes
}

// InfoNameVersion specializes the Information type as a list of named versions.
type InfoNameVersion []NameVersion

// Value returns the serialized form the of the list.
func (value InfoNameVersion) Value() []byte {
	var rawStruct C.cl_name_version
	rawStructSize := unsafe.Sizeof(rawStruct)
	bytes := make([]byte, rawStructSize*uintptr(len(value)))
	structOffset := uintptr(0)
	versionOffset := unsafe.Offsetof(rawStruct.version)
	nameOffset := unsafe.Offsetof(rawStruct.name)
	for _, entry := range value {
		copy(bytes[structOffset+nameOffset:structOffset+nameOffset+NameVersionNameMaxByteLength], []byte(entry.Name()))
		hostByteOrder.PutUint32(bytes[structOffset+versionOffset:structOffset+versionOffset+4], uint32(entry.Version))
		structOffset += rawStructSize
	}
	return bytes
}

var platformInfoDefaults = map[uint32]Information{
	C.CL_PLATFORM_PROFILE:    InfoString("FULL_PROFILE"),
	C.CL_PLATFORM_NAME:       InfoString("runtime-simulator"),
	C.CL_PLATFORM_VENDOR:     InfoString("github.com/opencl-go"),
	C.CL_PLATFORM_EXTENSIONS: InfoString("cl_khr_icd"),
	C.CL_PLATFORM_EXTENSIONS_WITH_VERSION: InfoNameVersion{
		NameVersion{Version: VersionOf(1, 0, 0), name: "cl_khr_icd"},
	},
	C.CL_PLATFORM_HOST_TIMER_RESOLUTION: InfoUint64(0),
	C.CL_PLATFORM_VERSION:               InfoString("OpenCL 3.0 "),
	C.CL_PLATFORM_NUMERIC_VERSION:       InfoVersion(VersionOf(3, 0, 0)),
}

// NewPlatform returns a new instance.
func NewPlatform() *Platform {
	platform := Platform{}
	return &platform
}

// Info returns the raw byte value for the queried information.
func (p *Platform) Info(infoName uint32) ([]byte, bool) {
	info, known := platformInfoDefaults[infoName]
	if !known {
		return nil, false
	}
	return info.Value(), true
}
