package airport

import (
	"context"
	"ronin/internal/model"
	"ronin/internal/repository"
)

// The Controller interface provides specification related to order functionality.
type Controller interface {
	RetrieveByCode(ctx context.Context, code string) (model.Airport, error)
}

type impl struct {
	repo repository.Registry
}

// New returns an implementation instance satisfying controller impl
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}
