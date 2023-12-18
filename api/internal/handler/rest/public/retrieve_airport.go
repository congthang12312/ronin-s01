package public

import (
	"errors"
	"net/http"
	"ronin/internal/handler/rest/public/response"
	"strings"
)

// RetrieveAirports retrieve a list of airport
func (h Handler) RetrieveAirports() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		code := strings.TrimSpace(r.URL.Query().Get("code"))
		if code == "" {
			response.ResponseError(ctx, w, errors.New("missing param"), http.StatusBadRequest)
			return
		}

		airport, err := h.airportCtrl.RetrieveByCode(ctx, code)
		if err != nil {
			response.ResponseError(ctx, w, err, http.StatusNotFound)
			return
		}

		response.RespondJSON(ctx, w, airportResponse{
			Code: airport.Code,
			Name: airport.Name,
		})
		return
	}
}

type airportResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
