package airport

import (
	"errors"
	"ronin/internal/model"
	"time"
)

// GetByCodeName return a list of Airport
func (i impl) GetByCodeName(code string) (model.Airport, error) {
	time.Sleep(4 * time.Second)
	if v, ok := airportMap[code]; ok {
		return v, nil
	}
	return model.Airport{}, errors.New("not found")
}
