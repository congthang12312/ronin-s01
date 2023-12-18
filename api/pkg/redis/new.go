package redis

import (
	"context"
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

// NewClient connects to Redis and returns a client which is used to access Redis
func NewClient(ctx context.Context, url string, database int, opts ...Option) (*Client, error) {

	pool := &redigo.Pool{
		MaxActive:       10,
		MaxIdle:         3,
		MaxConnLifetime: 9 * time.Minute, // https://docs.microsoft.com/en-us/azure/azure-cache-for-redis/cache-best-practices-connection#idle-timeout
		DialContext: func(ctx context.Context) (redigo.Conn, error) {
			return redigo.DialURLContext(
				ctx,
				url,
				redigo.DialDatabase(database),
			)
		},
	}

	cfg := config{
		pool: pool,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	client := Client{
		pool: pool,
	}

	if cfg.pingUponInit {
		if err := client.Ping(ctx); err != nil {
			return nil, err
		}
	}

	fmt.Println("Redis initialized")
	return &client, nil
}
