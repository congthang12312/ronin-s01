package redis

import (
	"context"
	"errors"

	redigo "github.com/gomodule/redigo/redis"
	pkgerrors "github.com/pkg/errors"
)

// Client is the default redis client
type Client struct {
	pool *redigo.Pool
}

// Close closes the connections in the pool
func (c *Client) Close() error {
	return c.pool.Close()
}

// Ping pings the redis instance
func (c *Client) Ping(ctx context.Context) error {
	var err error

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	defer conn.Close()

	p, err := redigo.DoContext(conn, ctx, "PING")
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	if p != "PONG" {
		err = pkgerrors.WithStack(errors.New("PONG not returned"))
		return err
	}
	return nil
}
