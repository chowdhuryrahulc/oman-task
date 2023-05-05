package util

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/chowdhuryrahulc/oman-task/src/main/dto/response"
)

// Send ...
func Send(w http.ResponseWriter, r *http.Request, message string, data interface{}) {
	if len(message) == 0 {
		message = "Success"
	}

	okRes := response.APIResponse{
		Status:        "ok",
		StatusCode:    0,
		StatusMessage: message,
		Data:          data,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, okRes)
}

// SendBadRequest ...
func SendBadRequest(w http.ResponseWriter, r *http.Request, message string) {
	if len(message) == 0 {
		message = "Bad Request"
	}

	badReqRes := response.APIResponse{
		Status:        "error",
		StatusCode:    1,
		StatusMessage: message,
	}

	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, badReqRes)
}

// SendForbidden ...
func SendForbidden(w http.ResponseWriter, r *http.Request, message string) {
	if len(message) == 0 {
		message = "Forbidden"
	}

	forbiddenRes := response.APIResponse{
		Status:        "error",
		StatusCode:    1,
		StatusMessage: message,
	}

	render.Status(r, http.StatusForbidden)
	render.JSON(w, r, forbiddenRes)
}

// SendInternalError ...
func SendInternalError(w http.ResponseWriter, r *http.Request, message string, err error) {
	if len(message) == 0 {
		message = "Internal Error"
	}

	internalErrRes := response.APIResponse{
		Status:        "error",
		StatusCode:    1,
		StatusMessage: message,
	}

	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, internalErrRes)
}

// SendNotFound ...
func SendNotFound(w http.ResponseWriter, r *http.Request, message string) {
	if len(message) == 0 {
		message = "Not Found"
	}

	notFoundRes := response.APIResponse{
		Status:        "error",
		StatusCode:    1,
		StatusMessage: message,
	}

	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, notFoundRes)
}
