package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse is ...
type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

var (
	// ErrMethodNotAllowed is ...
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
	// ErrNotFound is ...
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	// ErrBadRequest is ...
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
)

// Render is ...
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

// ErrorRenderer is ...
func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: "Bad request",
		Message:    err.Error(),
	}
}

// ServerErrorRenderer is ...
func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 500,
		StatusText: "Internal server error",
		Message:    err.Error(),
	}
}
