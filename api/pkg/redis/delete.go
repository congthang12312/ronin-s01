package redis

import (
	"context"
	"fmt"

	redigo "github.com/gomodule/redigo/redis"
	pkgerrors "github.com/pkg/errors"
)

// Del deletes the key and value from Redis
// keyObjectType is a prefix to denote the type of object the key will return.
// keyID is the variable part of the key.
// The actual redis key will be `keyObjectType:keyID` as suggested in https://redis.io/topics/data-types-intro
// eg - product:1 where bobcat is the projectName provided to NewClient, product is the keyObjectType and 1 is the keyID
func (c *Client) Del(ctx context.Context, keyObjectType, keyID string) error {
	key := fmt.Sprintf("%s:%s", keyObjectType, keyID)
	var err error

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	defer conn.Close()

	_, err = redigo.DoContext(conn, ctx, "DEL", key)
	err = pkgerrors.WithStack(err)
	return err
}
