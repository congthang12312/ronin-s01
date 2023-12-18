package redis

import (
	"context"
	"fmt"

	redigo "github.com/gomodule/redigo/redis"
	pkgerrors "github.com/pkg/errors"
)

// GetString gets a string value for the key from Redis
// keyObjectType is a prefix to denote the type of object the key will return.
// keyID is the variable part of the key.
// The actual redis key will be `keyObjectType:keyID` as suggested in https://redis.io/topics/data-types-intro.
// eg - product:1 where  product is the keyObjectType and 1 is the keyID
func (c *Client) GetString(ctx context.Context, keyObjectType, keyID string) (string, error) {
	r, err := c.get(ctx, keyObjectType, keyID) // Not doing redis.String directly to attach stacktrace for redis.String error
	if err != nil {
		return "", err
	}
	v, err := redigo.String(r, nil)
	return v, pkgerrors.WithStack(err)
}

// GetInt gets a string value for the key from Redis
// keyObjectType is a prefix to denote the type of object the key will return.
// keyID is the variable part of the key.
// The actual redis key will be `keyObjectType:keyID` as suggested in https://redis.io/topics/data-types-intro.
// eg - product:1 where product is the keyObjectType and 1 is the keyID
func (c *Client) GetInt(ctx context.Context, keyObjectType, keyID string) (int, error) {
	r, err := c.get(ctx, keyObjectType, keyID) // Not doing redis.Int directly to attach stacktrace for redis.Int error
	if err != nil {
		return 0, err
	}
	v, err := redigo.Int(r, nil)
	return v, pkgerrors.WithStack(err)
}

// GetInt64 gets a int64 value for the key from Redis
// keyObjectType is a prefix to denote the type of object the key will return.
// keyID is the variable part of the key.
// The actual redis key will be `keyObjectType:keyID` as suggested in https://redis.io/topics/data-types-intro.
// eg - product:1 where product is the keyObjectType and 1 is the keyID
func (c *Client) GetInt64(ctx context.Context, keyObjectType, keyID string) (int64, error) {
	r, err := c.get(ctx, keyObjectType, keyID) // Not doing redis.Int64 directly to attach stacktrace for redis.Int64 error
	if err != nil {
		return 0, err
	}
	v, err := redigo.Int64(r, nil)
	return v, pkgerrors.WithStack(err)
}

// Get gets a value by the key from Redis
// keyObjectType is a prefix to denote the type of object the key will return.
// keyID is the variable part of the key.
// The actual redis key will be `keyObjectType:keyID` as suggested in https://redis.io/topics/data-types-intro.
// eg - product:1 where product is the keyObjectType and 1 is the keyID
func (c *Client) get(ctx context.Context, keyObjectType, keyID string) (interface{}, error) {
	key := fmt.Sprintf("%s:%s", keyObjectType, keyID)
	var err error

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	defer conn.Close()

	val, err := redigo.DoContext(conn, ctx, "GET", key)
	err = pkgerrors.WithStack(err)
	return val, err
}
