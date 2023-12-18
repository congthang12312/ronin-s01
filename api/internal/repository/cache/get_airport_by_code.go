package cache

import (
	"context"
	"encoding/json"
	"errors"
	"ronin/internal/model"

	"github.com/gomodule/redigo/redis"
)

const (
	// KeyObjectTypeAirport is key of airport object type
	KeyObjectTypeAirport = "airport_code"
)

// GetAirportByCode get airport by code in redis cache
func (i impl) GetAirportByCode(ctx context.Context, code string) (model.Airport, error) {
	value, err := i.redisClient.GetString(ctx, KeyObjectTypeAirport, code)
	if errors.Is(errors.Unwrap(err), redis.ErrNil) {
		return model.Airport{}, ErrKeyNotFound
	}

	var airport model.Airport
	if err := json.Unmarshal([]byte(value), &airport); err != nil {
		return model.Airport{}, err
	}

	return airport, nil
}
