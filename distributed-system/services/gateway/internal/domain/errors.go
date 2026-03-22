package domain

import "errors"

var (
	ErrRepositoryNotFound = errors.New("repository not found")
	ErrInvalidArgument    = errors.New("owner and name must not be empty")
)
