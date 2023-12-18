package router

import (
	"net/http"
	publicREST "ronin/internal/handler/rest/public"

	"github.com/go-chi/chi/v5"
)

// Router defines the routes & handlers of the app
type Router struct {
	publicRESTHandler publicREST.Handler
}

// Handler returns the Handler for use by the server
func (rtr Router) Handler() http.Handler {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Group(rtr.routes)
	})
	return r
}

// routes is place list of routers: public, m2m, b2c, authenticated...
func (rtr Router) routes(r chi.Router) {
	r.Group(rtr.public)
}
