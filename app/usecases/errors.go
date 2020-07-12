package usecases

import "errors"

var (
	errUserAlreadyExist   = errors.New("user exists")
	errUserRegistration   = errors.New("is not possible register the user")
	errInvalidCredentials = errors.New("invalid login credentials")
	errUnauthorizedUser   = errors.New("unauthorized user")

	errGameNotPlayable = errors.New("the game is not playable")
)
