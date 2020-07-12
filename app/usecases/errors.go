package usecases

import "errors"

var (
	errUserAlreadyExist   = errors.New("user exists")
	errUserRegistration   = errors.New("is not possible register the user")
	ErrInvalidCredentials = errors.New("invalid login credentials. Please try again")

	errGameNotPlayable = errors.New("the game is not playable")
)
