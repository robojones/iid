package iid

import "github.com/pkg/errors"

// ErrInvalid is returned when trying to import an invalid Iid from a string.
var ErrInvalid = errors.New("invalid id")

// IsErrInvalid returns true if the cause of the provided error is ErrInvalid.
func IsErrInvalid(err error) bool {
	return errors.Cause(err) == ErrInvalid
}
