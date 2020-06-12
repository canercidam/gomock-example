package api

import "errors"

var (
	// ErrInvalidParam is returned when API refuses query due to
	// invalid param.
	ErrInvalidParam = errors.New("invalid param")
)

// Not implemented.
