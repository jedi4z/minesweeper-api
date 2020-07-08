package mysql_repositories

import (
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jedi4z/minesweeper-api/app/repositories"
	"github.com/jinzhu/gorm"
)

type CellRepository struct {
	DB *gorm.DB
}

func NewCellRepository(db *gorm.DB) repositories.CellRepositoryInterface {
	return CellRepository{
		DB: db,
	}
}

func (r CellRepository) Find(id uint) (*models.Cell, error) {
	cell := new(models.Cell)

	if err := r.DB.Find(cell, id).Error; err != nil {
		return nil, err
	}

	return cell, nil
}

func (r CellRepository) Update(cell *models.Cell) error {
	return r.DB.Model(&models.Cell{}).Updates(cell).Error
}
