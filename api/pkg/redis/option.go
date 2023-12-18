package redis

import (
	redigo "github.com/gomodule/redigo/redis"
)

type config struct {
	pool         *redigo.Pool
	pingUponInit bool
}

// Option is an optional config used to modify the client's behaviour
type Option func(*config)
