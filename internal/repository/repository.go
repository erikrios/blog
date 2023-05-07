package repository

import "errors"

var (
	ErrNotFound      = errors.New("repository: requested data not found")
	ErrAlreadyExists = errors.New("repository: data already exists")
	ErrUnknown       = errors.New("repository: unknown error")
)
