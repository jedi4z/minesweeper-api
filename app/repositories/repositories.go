package repositories

import "github.com/jedi4z/minesweeper-api/app/models"

type UserRepositoryInterface interface {
	Insert(user *models.User) error
	IsUnique(user *models.User) (bool, error)
	FindOne(id string) (*models.User, error)
	FindOneByEmail(email string) (*models.User, error)
}

type GameRepositoryInterface interface {
	Insert(game *models.Game) error
	FindAll() ([]models.Game, error)
	Find(id uint) (*models.Game, error)
	Update(cell *models.Game) error
}
