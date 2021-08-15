package rest

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rmar8138/article-rest-api/internal"
)

// ErrorResponse represents an error response sent to the client
type ErrorResponse struct {
	Error string `json:"error"`
}

func handleErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	resp := ErrorResponse{Error: err.Error()}
	status := http.StatusInternalServerError
	var ierr *internal.Error

	if !errors.As(err, &ierr) {
		resp.Error = "internal error"
	} else {
		switch ierr.Code() {
		case internal.ErrorCodeNotFound:
			status = http.StatusNotFound
		case internal.ErrorCodeInvalidArgument:
			status = http.StatusBadRequest
		case internal.ErrorCodeUnknown:
			fallthrough
		default:
			status = http.StatusInternalServerError
		}
	}
	render.Status(r, status)
	render.JSON(w, r, resp)
}
