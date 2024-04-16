package storage

import "errors"

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotExist = errors.New("user doesn't exist")
)
