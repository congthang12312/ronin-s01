package response

import (
	"fmt"
	"net/http"
)

const (
	// DefaultErrorCode represents the default code for Error
	DefaultErrorCode = "internal_error"
	// DefaultErrorDesc represents the default description for Error
	DefaultErrorDesc = "Something went wrong"
)

var (
	// ErrDefaultInternal represents the default internal server error
	ErrDefaultInternal = &Error{
		Status: http.StatusInternalServerError,
		Code:   DefaultErrorCode,
		Desc:   DefaultErrorDesc,
	}
)

// Error represents a handler error. It contains web-related information such as HTTP status code, error code and
// error description
type Error struct {
	Status int    `json:"-"`
	Code   string `json:"error"`             // Since there is existing dependency at infinity end, unable to fix the json key name
	Desc   string `json:"error_description"` // Since there is existing dependency at infinity end, unable to fix the json key name
}

// Error satisfies the error interface
func (e Error) Error() string {
	return fmt.Sprintf("Status: [%d], Code: [%s], Desc: [%s]", e.Status, e.Code, e.Desc)
}
