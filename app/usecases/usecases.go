package usecases

import "github.com/jedi4z/minesweeper-api/app/models"

type UserUseCaseInterface interface {
	RegisterUser(user *models.User) error
	AuthenticateUser(user *models.User) (*string, error)
	FindByAccessToken(accessToken string) (*models.User, error)
}

type GameUseCasesInterface interface {
	CreateGame(game *models.Game) error
	ListGames(user *models.User) ([]models.Game, error)
	GetGame(user *models.User, id uint) (*models.Game, error)
	HoldGame(game *models.Game) error
	ResumeGame(game *models.Game) error
	FlagCell(game *models.Game, cellID uint) error
	UncoverCell(game *models.Game, cellID uint) error
}
