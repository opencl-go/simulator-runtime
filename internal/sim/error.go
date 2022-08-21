package sim

// DomainError represents a basic error that occurs within the simulator domain.
type DomainError string

// String returns the error text.
func (err DomainError) Error() string {
	return string(err)
}
