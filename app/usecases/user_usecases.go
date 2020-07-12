// Package registration represents the concrete implementation of UserCaseInterface interface.
// Because the same business function can be created to support both transaction and non-transaction,
// a shared business function is created in a helper file, then we can wrap that function with transaction
// or non-transaction.
package usecases

import (
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jedi4z/minesweeper-api/app/repositories"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserUseCase(r repositories.UserRepositoryInterface) UserUseCaseInterface {
	return UserUseCase{
		UserRepository: r,
	}
}

func (uc UserUseCase) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.WithError(err).Error()
		return ErrUserRegistration
	}

	user.Password = string(hashedPassword)

	unique, err := uc.UserRepository.IsUnique(user)
	if unique == false {
		log.WithError(err).Error()
		return ErrUserAlreadyExist
	}

	if err := uc.UserRepository.Insert(user); err != nil {
		log.WithError(err).Error()
		return ErrUserRegistration
	}

	return nil
}

func (uc UserUseCase) AuthenticateUser(email string, password string) (*string, error) {
	panic("not implemented")
}
