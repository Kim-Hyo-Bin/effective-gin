package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRequest     = New("invalid request")
	ErrUserNotFound       = New("user not found")
	ErrUnauthorized       = New("unauthorized access")
	ErrInternalServer     = New("internal server error")
	ErrExternalAPI        = New("external API failure")
	ErrInvalidInput       = New("invalid input")
	ErrDatabaseError      = New("database error")
	ErrFileNotFound       = New("file not found")
	ErrTimeout            = New("timeout")
	ErrConflict           = New("conflict")
	ErrBadGateway         = New("bad gateway")
	ErrServiceUnavailable = New("service unavailable")
	ErrGatewayTimeout     = New("gateway timeout")
	ErrPermissionDenied   = New("permission denied")
	ErrRateLimitExceeded  = New("rate limit exceeded")
	ErrInvalidToken       = New("invalid token")
	ErrResourceExhausted  = New("resource exhausted")

	// ErrNotImplemented is a placeholder for not implemented errors.
	ErrNotImplemented = New("not implemented")
)

func New(baseErr string) error {
	return errors.New(baseErr)
}

func Wrap(baseErr string, message string) error {
	return fmt.Errorf("%s: %w", message, New(baseErr))
}

func WithMessage(err error, message string) error {
	return fmt.Errorf("%s: %v", message, err)
}

func Cause(err error) error {
	for {
		unwrappedErr := errors.Unwrap(err)
		if unwrappedErr == nil {
			return err
		}
		err = unwrappedErr
	}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}
