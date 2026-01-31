package domain

import "errors"

var (
	// User Domain Errors
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidUserEmail   = errors.New("invalid user email")
	ErrDeletedUserAccess   = errors.New("access to deleted user")
	ErrInvalidUserId	   = errors.New("invalid user ID")
)