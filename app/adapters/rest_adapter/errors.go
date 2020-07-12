package rest_adapter

import "errors"

var (
	errCredentialRequired = errors.New("credential required")
	errInvalidAccessToken = errors.New("invalid access token")
	errUserUnauthorized   = errors.New("user unauthorized")
)
