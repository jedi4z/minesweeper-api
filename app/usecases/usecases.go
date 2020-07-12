package usecases

import "github.com/jedi4z/minesweeper-api/app/models"

type UserUseCaseInterface interface {
	RegisterUser(user *models.User) error
	AuthenticateUser(email string, password string) (*string, error)
}

type GameUseCasesInterface interface {
	CreateGame(game *models.Game) error
	ListGames() ([]models.Game, error)
	GetGame(id uint) (*models.Game, error)
	HoldGame(id uint) (*models.Game, error)
	ResumeGame(id uint) (*models.Game, error)
	FlagCell(game *models.Game, cellID uint) error
	UncoverCell(game *models.Game, cellID uint) error
}
