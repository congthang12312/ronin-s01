package redis

import (
	"errors"
)

var (
	// ErrSetFailed means the set command did not return an OK response
	ErrSetFailed = errors.New("set command failed")
)
