package client

// Error constants for client operations.
const (
	// ErrInvalidClientMACFormat is the error message for invalid client MAC address format.
	ErrInvalidClientMACFormat = "invalid client MAC address: %s"

	// ErrEmptyClientIPAddr is the error message for an empty client IP address.
	ErrEmptyClientIPAddr = "client IP address cannot be empty"

	// ErrEmptyClientUsername is the error message for an empty client username.
	ErrEmptyClientUsername = "client username cannot be empty"
)
