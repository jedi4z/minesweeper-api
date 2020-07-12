package usecases

import "errors"

var (
	errUserAlreadyExist = errors.New("user exists")
	errUserRegistration = errors.New("is not possible register the user")

	errGameNotPlayable = errors.New("the game is not playable")
)
