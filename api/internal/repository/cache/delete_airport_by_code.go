package cache

import "context"

// DeleteAirportByCode delete code from cache
func (i impl) DeleteAirportByCode(ctx context.Context, code string) error {
	return i.redisClient.Del(ctx, KeyObjectTypeAirport, code)
}
