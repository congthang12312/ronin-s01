package router

import "github.com/go-chi/chi/v5"

func (rtr Router) public(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Get("/airport-service/v1/airports", rtr.publicRESTHandler.RetrieveAirports())
	})
}
