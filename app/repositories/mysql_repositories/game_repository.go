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

func (r GameRepository) FindAll() ([]models.Game, error) {
	games := make([]models.Game, 0)

	if err := r.DB.Find(&games).Error; err != nil {
		return nil, err
	}

	return games, nil

}

func (r GameRepository) Find(id uint) (*models.Game, error) {
	game := new(models.Game)

	if err := r.DB.
		Preload("Grid").
		Preload("Grid.Cells").
		Find(game, id).Error; err != nil {
		return nil, err
	}

	return game, nil
}

func (r GameRepository) FindOneByCellID(cellID uint) (*models.Game, error) {
	game := new(models.Game)

	if err := r.DB.
		Preload("Grid").
		Preload("Grid.Cells").
		Joins("JOIN `rows` AS r ON r.game_id = games.id").
		Joins("JOIN `cells` AS c ON c.row_id = r.id").
		Where("c.id = ?", cellID).
		First(game).Error; err != nil {
		return nil, err
	}

	return game, nil
}
