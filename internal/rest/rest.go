package rest

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	Err string `json:"error"`
}

func handleErrorBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, errorResponse{
		Err: err.Error(),
	})
}
