package repositories

import "github.com/jedi4z/minesweeper-api/app/models"

type GameRepositoryInterface interface {
	Insert(game *models.Game) error
	FindAll() ([]models.Game, error)
	Find(id uint) (*models.Game, error)
	Update(cell *models.Game) error
}
