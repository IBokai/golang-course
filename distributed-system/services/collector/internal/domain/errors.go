package domain

import "errors"

var (
	ErrRepositoryNotFound = errors.New("repository not found")
	ErrRateLimitExceeded  = errors.New("github rate limit exceeded")
	ErrInvalidArgument    = errors.New("owner and name must not be empty")
)
