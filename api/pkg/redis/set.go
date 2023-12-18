package redis

import (
	"context"
	"fmt"
	"strconv"

	redigo "github.com/gomodule/redigo/redis"
	pkgerrors "github.com/pkg/errors"
)

// Set sets a key/value pair which expires in Redis. If key already exists, it will override the value.
// All keys must have expiration
// keyObjectType is a prefix to denote the type of object the key will return.
// keyID is the variable part of the key.
// The actual redis key will be `keyObjectType:keyID` as suggested in https://redis.io/topics/data-types-intro
// eg - product:1 where product is the keyObjectType and 1 is the keyID
func (c *Client) Set(
	ctx context.Context,
	keyObjectType string,
	keyID string,
	value interface{},
	expirationInSeconds int,
) error {
	key := fmt.Sprintf("%s:%s", keyObjectType, keyID)
	var err error

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	defer conn.Close()

	v, err := redigo.String(redigo.DoContext(conn, ctx, "SET", key, value, "EX", strconv.Itoa(expirationInSeconds)))
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	if v != "OK" {
		err = pkgerrors.WithStack(ErrSetFailed)
		return err
	}
	return nil
}

// SetIfNotExists sets a key/value pair ONLY if the key does not already exist. It also sets an expiry to this data.
// All keys must have expiration
// keyObjectType is a prefix to denote the type of object the key will return.
// keyID is the variable part of the key.
// The actual redis key will be `keyObjectType:keyID` as suggested in https://redis.io/topics/data-types-intro
// eg - product:1 where  product is the keyObjectType and 1 is the keyID
func (c *Client) SetIfNotExists(
	ctx context.Context,
	keyObjectType string,
	keyID string,
	value interface{},
	expirationInSeconds int,
) error {
	key := fmt.Sprintf("%s:%s", keyObjectType, keyID)
	var err error

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	defer conn.Close()

	v, err := redigo.String(redigo.DoContext(conn, ctx, "SET", key, value, "NX", "EX", strconv.Itoa(expirationInSeconds)))
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	if v != "OK" {
		err = pkgerrors.WithStack(ErrSetFailed)
		return err
	}
	return nil
}
