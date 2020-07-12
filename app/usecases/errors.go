package usecases

import "errors"

var (
	ErrUserAlreadyExist = errors.New("user exists")
	ErrUserRegistration = errors.New("is not possible register the user")

	ErrGameNotPlayable = errors.New("the game is not playable")
)
