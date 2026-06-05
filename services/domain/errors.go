package domain

import (
	"errors"
)

var ErrInternal = errors.New("internal error")
var ErrNotFound = errors.New("not found")
var ErrUnique = errors.New("already exists")
var ErrNotAllowed = errors.New("not allowed")
var ErrInvalidState = errors.New("invalid state")
var ErrStale = errors.New("stale data")
