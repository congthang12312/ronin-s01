package public

import "ronin/internal/controller/airport"

// Handler is the web handler for this pkg
type Handler struct {
	airportCtrl airport.Controller
}

// New instantiates a new Handler and returns it
func New(airportCtrl airport.Controller) Handler {
	return Handler{
		airportCtrl: airportCtrl,
	}
}
