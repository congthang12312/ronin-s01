package cache

import (
	"context"
	"ronin/internal/model"
	"ronin/pkg/redis"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	SetAirportByCode(ctx context.Context, code string, airport model.Airport, expirationInSeconds int) error
	GetAirportByCode(ctx context.Context, code string) (model.Airport, error)
	DeleteAirportByCode(ctx context.Context, code string) error
}

// New returns an implementation instance satisfying Repository
func New(redisClient *redis.Client) Repository {
	return impl{redisClient: redisClient}
}

type impl struct {
	redisClient *redis.Client
}
