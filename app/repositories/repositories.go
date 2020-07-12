package repositories

import "github.com/jedi4z/minesweeper-api/app/models"

type UserRepositoryInterface interface {
	Insert(user *models.User) error
	IsUnique(user *models.User) (bool, error)
	FindOne(id uint) (*models.User, error)
	FindOneByEmail(email string) (*models.User, error)
}

type GameRepositoryInterface interface {
	Insert(game *models.Game) error
	FindAll(user *models.User) ([]models.Game, error)
	Find(user *models.User, id uint) (*models.Game, error)
	Update(cell *models.Game) error
}
