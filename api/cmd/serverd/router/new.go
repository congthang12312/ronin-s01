package router

import (
	"ronin/internal/controller/airport"
	publicREST "ronin/internal/handler/rest/public"
)

// New creates and returns a new Router instance
func New(
	airportCtrl airport.Controller,
) Router {
	return Router{
		publicRESTHandler: publicREST.New(airportCtrl),
	}
}
