package usecases

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jedi4z/minesweeper-api/app/repositories"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const secret = "this-should-be-in-a-secure-place"

type UserUseCase struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserUseCase(r repositories.UserRepositoryInterface) UserUseCaseInterface {
	return UserUseCase{
		UserRepository: r,
	}
}

func (uc UserUseCase) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
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
	userDB, err := uc.UserRepository.FindOneByEmail(user.Email)
	if err != nil {
		log.WithError(err).Error()
		return nil, errInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.WithError(err).Error()
		return nil, errInvalidCredentials
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	claims := &models.Claims{
		ID:    userDB.ID,
		Email: userDB.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(secret))
	if err != nil {
		log.WithError(err).Error()
		return nil, errInvalidCredentials
	}

	return &token, nil
}

func verifyAccessToken(at string) (*models.Claims, error) {
	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(at, claims, func(tk *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !tkn.Valid {
		log.WithError(err).Error()
		return nil, errUnauthorizedUser
	}

	return claims, nil
}

func (uc UserUseCase) FindByAccessToken(accessToken string) (*models.User, error) {
	claims, err := verifyAccessToken(accessToken)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	// retrieve the user from database to check if exists
	return uc.UserRepository.FindOne(claims.ID)
}
