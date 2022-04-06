package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method no allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Resource not found"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrorRenderer(err error) *ErrorResponse {
	return errorRendererFactory(400, "Bad Request", err)
}

func ServerErrorRenderer(err error) *ErrorResponse {
	return errorRendererFactory(500, "Internal server error", err)
}

func errorRendererFactory(statusCode int, statusText string, err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: statusCode,
		StatusText: statusText,
		Message:    err.Error(),
	}

}
