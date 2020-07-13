package rest_adapter

import (
	"net/http"
)

type restError struct {
	StatusCode int    `json:"status_code,omitempty"`
	ErrorCode  string `json:"error_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func newUnauthorizedRestError(err error) restError {
	return restError{
		StatusCode: http.StatusUnauthorized,
		ErrorCode:  "unauthorized",
		Message:    err.Error(),
	}
}

func newBadRequestRestError(err error) restError {
	return restError{
		StatusCode: http.StatusBadRequest,
		ErrorCode:  "bad_request",
		Message:    err.Error(),
	}
}

func newNotFoundRestError(err error) restError {
	return restError{
		StatusCode: http.StatusNotFound,
		ErrorCode:  "not_found",
		Message:    err.Error(),
	}
}

func newInternalServerRestError(err error) restError {
	return restError{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  "internal_server_error",
		Message:    err.Error(),
	}
}

func newForbiddenRestError(err error) restError {
	return restError{
		StatusCode: http.StatusForbidden,
		ErrorCode:  "forbidden",
		Message:    err.Error(),
	}
}
