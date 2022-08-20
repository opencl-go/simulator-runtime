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

var platformInfoDefaults = map[uint32]Information{
	C.CL_PLATFORM_PROFILE:                 InfoString("FULL_PROFILE"),
	C.CL_PLATFORM_NAME:                    InfoString("runtime-simulator"),
	C.CL_PLATFORM_VENDOR:                  InfoString("github.com/opencl-go"),
	C.CL_PLATFORM_EXTENSIONS:              InfoString(""),
	C.CL_PLATFORM_EXTENSIONS_WITH_VERSION: InfoZero{},
	C.CL_PLATFORM_HOST_TIMER_RESOLUTION:   InfoUint64(0),
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
