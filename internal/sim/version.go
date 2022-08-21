package sim

// #include "../api/api.h"
import "C"
import (
	"fmt"
	"math"
)

// Version represents a major.minor.patch version combination, encoded in a 32-bit unsigned integer value.
type Version uint32

const (
	// VersionMin is the lowest value that Version can represent. It has all components set to zero.
	VersionMin Version = 0
	// VersionMax is the highest value that Version can represent.
	// It has all components set to their maximally possible value.
	VersionMax Version = math.MaxUint32

	versionMajorBits = 10
	versionMinorBits = 10
	versionPatchBits = 12
	versionMajorMask = (1 << versionMajorBits) - 1
	versionMinorMask = (1 << versionMinorBits) - 1
	versionPatchMask = (1 << versionPatchBits) - 1
)

// VersionOf returns a Version that has the three provided components encoded.
// No particular limits-checking is performed, the provided values are cast and shifted into the final field.
func VersionOf(major, minor, patch int) Version {
	return Version(
		((uint32(major) & versionMajorMask) << (versionMinorBits + versionPatchBits)) |
			((uint32(minor) & versionMinorMask) << versionPatchBits) | (uint32(patch) & versionPatchMask))
}

// String returns a common representation of <major> <.> <minor> <.> <patch> of the version, with the components
// in decimal format.
func (ver Version) String() string {
	return fmt.Sprintf("%d.%d.%d", ver.Major(), ver.Minor(), ver.Patch())
}

// Major returns the major component value.
func (ver Version) Major() int {
	return int((uint32(ver) >> (versionMinorBits + versionPatchBits)) & versionMajorMask)
}

// Minor returns the minor component value.
func (ver Version) Minor() int {
	return int((uint32(ver) >> versionPatchBits) & versionMinorMask)
}

// Patch returns the patch component value.
func (ver Version) Patch() int {
	return int(uint32(ver) & versionPatchMask)
}

// NameVersion is a combination of a name with a version.
type NameVersion struct {
	// Version describes the maturity level of the identified element.
	Version Version
	name    string
}

const (
	// NameVersionNameMaxByteLength is the maximum length, in bytes, that a name component of NameVersion may have.
	NameVersionNameMaxByteLength = C.CL_NAME_VERSION_MAX_NAME_SIZE - 1

	// ErrNameTooLong is returned from NameVersionFrom() if the byte-length of the provided value is longer than OpenCL supports.
	// Refer to NameVersionNameMaxByteLength for the limit.
	ErrNameTooLong = DomainError("name for named version is too long")

	// ErrNameEmpty is reported if a NameVersion was found to have no name.
	ErrNameEmpty = DomainError("name must not be empty")
)

// NameVersionFrom returns an instance, if the given parameters are valid.
func NameVersionFrom(name string, version Version) (NameVersion, error) {
	bytesLen := len([]byte(name))
	if bytesLen == 0 {
		return NameVersion{}, ErrNameEmpty
	} else if bytesLen > NameVersionNameMaxByteLength {
		return NameVersion{}, ErrNameTooLong
	}
	return NameVersion{Version: version, name: name}, nil
}

// Name returns the text value of the versioned component.
func (ver NameVersion) Name() string {
	return ver.name
}
