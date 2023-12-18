package response

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Success is the response format when http handler succeeds
type Success struct {
	Message string `json:"message,omitempty"`
}

// RespondJSON handles conversion of the requested result to JSON format
func RespondJSON(ctx context.Context, w http.ResponseWriter, obj interface{}) {
	RespondJSONWithHeaders(ctx, w, obj, nil)
}

// RespondJSONWithHeaders handles conversion of the requested result to JSON format
func RespondJSONWithHeaders(ctx context.Context, w http.ResponseWriter, obj interface{}, headers map[string]string) {
	var respBytes []byte
	var err error

	// Set HTTP headers
	w.Header().Set("Content-Type", "application/json")
	for key, value := range headers {
		w.Header().Set(key, value)
	}

	status := http.StatusOK
	switch parsed := obj.(type) {
	case *Error:
		if parsed.Status >= http.StatusInternalServerError && parsed.Status != http.StatusServiceUnavailable {
			parsed.Desc = DefaultErrorDesc
		}
		status = parsed.Status
		if status == 0 {
			status = http.StatusInternalServerError
		}
		respBytes, err = json.Marshal(parsed)
	case error:
		status = http.StatusInternalServerError
		respBytes, err = json.Marshal(ErrDefaultInternal)
	default:
		respBytes, err = json.Marshal(obj)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(status)

	if _, err = w.Write(respBytes); err != nil {
		fmt.Print("[RespondJSON] write failed")
	}
}

// ResponseError return http error
func ResponseError(ctx context.Context, w http.ResponseWriter, err error, statusCode int) {
	RespondJSON(ctx, w, &Error{
		Status: statusCode,
		Desc:   err.Error(),
	})
}
