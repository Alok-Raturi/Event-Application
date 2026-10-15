package user

import "errors"

var (
	ErrInvalidInput        = errors.New("invalid input provided by user")
	ErrEmailAlreadyExists  = errors.New("user with this email already exists")
	ErrJsonEncodingError   = errors.New("Failure in encoding the response")
	ErrResponseClientError = errors.New("Failure in writing data in response")
)
