package storage

import "errors"

var (
	ErrUserCreateFailed = errors.New("failed to create user")
	ErrUserNotFound     = errors.New("user not found")
	ErrInternalError    = errors.New("internal error")
)
