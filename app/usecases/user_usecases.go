// Package registration represents the concrete implementation of UserCaseInterface interface.
// Because the same business function can be created to support both transaction and non-transaction,
// a shared business function is created in a helper file, then we can wrap that function with transaction
// or non-transaction.
package usecases

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jedi4z/minesweeper-api/app/repositories"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
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
		return errUserRegistration
	}

	user.Password = string(hashedPassword)

	unique, err := uc.UserRepository.IsUnique(user)
	if unique == false {
		log.WithError(err).Error()
		return errUserAlreadyExist
	}

	if err := uc.UserRepository.Insert(user); err != nil {
		log.WithError(err).Error()
		return errUserRegistration
	}

	return nil
}

func (uc UserUseCase) AuthenticateUser(user *models.User) (*string, error) {
	email, password := user.Email, user.Password

	u, err := uc.UserRepository.FindOneByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.WithError(err).Error()
		return nil, ErrInvalidCredentials
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	claims := &models.Claims{
		ID:    u.CommonFields.ID,
		Email: u.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := "this-should-be-in-a-secure-place"
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		log.WithError(err).Error()
		return nil, ErrInvalidCredentials
	}

	return &token, nil
}
