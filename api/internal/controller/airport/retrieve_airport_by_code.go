package airport

import (
	"context"
	"errors"
	"ronin/internal/model"
	"ronin/internal/repository/cache"
)

// airportExpiryTime expiry time of  airport within 10s
const airportExpiryTime = 10

// RetrieveByCode retrieve Airport by code
func (i impl) RetrieveByCode(ctx context.Context, code string) (model.Airport, error) {
	airportCache, err := i.repo.Cache().GetAirportByCode(ctx, code)
	if err != nil && !errors.Is(err, cache.ErrKeyNotFound) {
		return model.Airport{}, err
	}

	if airportCache.Code != "" {
		return airportCache, nil
	}

	airport, err := i.repo.Airport().GetByCodeName(code)
	if err != nil {
		return model.Airport{}, err
	}

	if err := i.repo.Cache().SetAirportByCode(ctx, code, airport, airportExpiryTime); err != nil {
		return model.Airport{}, err
	}

	return airport, nil
}
