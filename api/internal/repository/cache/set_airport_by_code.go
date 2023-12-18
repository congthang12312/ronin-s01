package cache

import (
	"context"
	"encoding/json"
	"ronin/internal/model"
)

// SetAirportByCode set the airport to cache
func (i impl) SetAirportByCode(ctx context.Context, code string, airport model.Airport, expirationInSeconds int) error {
	value, err := json.Marshal(airport)
	if err != nil {
		return err
	}
	return i.redisClient.Set(ctx, KeyObjectTypeAirport, code, value, expirationInSeconds)
}
