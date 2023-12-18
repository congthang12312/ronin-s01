package repository

import (
	"ronin/internal/repository/airport"
	"ronin/internal/repository/cache"
	"ronin/pkg/redis"
)

// Registry is the registry of all the domain specific repositories and also provides transaction capabilities.
type Registry interface {
	// Airport returns the airport repo
	Airport() airport.Repository

	// Cache return the redis repo
	Cache() cache.Repository
}

// New returns a new instance of Registry
func New(redisClient *redis.Client) Registry {
	return impl{
		redisClient: redisClient,
	}
}

type impl struct {
	redisClient *redis.Client
}

// Airport return the airport repo
func (i impl) Airport() airport.Repository {
	return airport.New(i.redisClient)
}

// Cache returns the redis repo
func (i impl) Cache() cache.Repository {
	return cache.New(i.redisClient)
}
