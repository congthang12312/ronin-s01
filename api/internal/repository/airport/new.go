package airport

import (
	"ronin/internal/model"
	"ronin/pkg/redis"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	GetByCodeName(code string) (model.Airport, error)
}

// New returns an implementation instance satisfying Repository
func New(redisClient *redis.Client) Repository {
	return impl{redisClient: redisClient}
}

type impl struct {
	redisClient *redis.Client
}
