package sim

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
