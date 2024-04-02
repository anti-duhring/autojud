package user

import "errors"

var (
	ErrNameEmpty          = errors.New("Name can't be empty")
	ErrEmailEmpty         = errors.New("Email can't be empty")
	ErrPasswordEmpty      = errors.New("Password can't be empty")
	ErrEmailInvalid       = errors.New("Email is invalid")
	ErrInternal           = errors.New("Internal error")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)
