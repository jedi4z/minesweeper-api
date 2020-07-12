package mysql_repositories

import (
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jedi4z/minesweeper-api/app/repositories"
	"github.com/jinzhu/gorm"
)

type GameRepository struct {
	DB *gorm.DB
}

func NewGameRepository(db *gorm.DB) repositories.GameRepositoryInterface {
	return GameRepository{
		DB: db,
	}
}

func (r GameRepository) Insert(game *models.Game) error {
	return r.DB.Create(&game).Error
}

func (r GameRepository) FindAll(user *models.User) ([]models.Game, error) {
	games := make([]models.Game, 0)

	if err := r.DB.
		Where("user_id = ?", user.ID).
		Find(&games).Error; err != nil {
		return nil, err
	}

	return games, nil

}

func (r GameRepository) Find(user *models.User, id uint) (*models.Game, error) {
	game := new(models.Game)

	if err := r.DB.
		Preload("Grid").
		Preload("Grid.Cells").
		Where("user_id = ?", user.ID).
		Find(game, id).Error; err != nil {
		return nil, err
	}

	return game, nil
}

func (r GameRepository) Update(game *models.Game) error {
	return r.DB.Model(&models.Game{}).Updates(game).Error
}
