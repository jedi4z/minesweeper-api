package usecases

import "github.com/jedi4z/minesweeper-api/app/models"

type GameUseCasesInterface interface {
	CreateGame(game *models.Game) error
	ListGames() ([]models.Game, error)
	GetGame(id uint) (*models.Game, error)
	GetGameByCellID(cellID uint) (*models.Game, error)
}

type CellUseCasesInterface interface {
	GetCell(id uint) (*models.Cell, error)
	ClickCell(cell *models.Cell, game *models.Game) error
}
