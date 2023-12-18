package airport

import (
	"ronin/internal/model"
)

var (
	airportMap = map[string]model.Airport{
		"VJ1": model.Airport{
			Code: "VJ1",
			Name: "Vietjet 1",
		},
		"VJ2": model.Airport{
			Code: "VJ2",
			Name: "Vietjet 2",
		},
		"VJ3": model.Airport{
			Code: "VJ3",
			Name: "Vietjet 3",
		},
		"VJ4": model.Airport{
			Code: "V4",
			Name: "Vietjet 4",
		},
	}
)
