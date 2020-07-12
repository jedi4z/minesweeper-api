package rest_adapter

import (
	"net/http"
)

type restError struct {
	StatusCode int    `json:"status,omitempty"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func newUnauthorizedRestError(err error) restError {
	return restError{
		StatusCode: http.StatusUnauthorized,
		Code:       "unauthorized",
		Message:    err.Error(),
	}
}

func newBadRequestRestError(err error) restError {
	return restError{
		StatusCode: http.StatusBadRequest,
		Code:       "bad_request",
		Message:    err.Error(),
	}
}

func newNotFoundRestError(err error) restError {
	return restError{
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
		Message:    err.Error(),
	}
}

func newInternalServerRestError(err error) restError {
	return restError{
		StatusCode: http.StatusInternalServerError,
		Code:       "internal_server_error",
		Message:    err.Error(),
	}
}

func newForbiddenRestError(err error) restError {
	return restError{
		StatusCode: http.StatusForbidden,
		Code:       "forbidden",
		Message:    err.Error(),
	}
}
