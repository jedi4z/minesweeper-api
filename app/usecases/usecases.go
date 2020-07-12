package usecases

import "github.com/jedi4z/minesweeper-api/app/models"

type GameUseCasesInterface interface {
	CreateGame(game *models.Game) error
	ListGames() ([]models.Game, error)
	GetGame(id uint) (*models.Game, error)
	UncoverCell(game *models.Game, cellID uint) error
}
