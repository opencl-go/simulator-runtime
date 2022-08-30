package sim

import "fmt"

// APIObjectAllocator creates new instances of API objects.
type APIObjectAllocator interface {
	New(ref any) APIObject
}

// APIObject represents some object available on the API.
type APIObject interface {
	Delete()
	ID() ObjectID
}

// ObjectID represents a unique identifier of an API object.
type ObjectID uintptr

// String returns the string presentation of the underlying uintptr value.
func (id ObjectID) String() string {
	return fmt.Sprintf("0x%X", uintptr(id))
}
