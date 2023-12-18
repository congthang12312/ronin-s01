package cache

import "errors"

var (
	// ErrKeyNotFound is when key not found
	ErrKeyNotFound = errors.New("key not found")
)
