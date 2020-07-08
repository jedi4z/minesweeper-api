package usecases

import (
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jedi4z/minesweeper-api/app/repositories"
)

type CellUseCases struct {
	CellRepository repositories.CellRepositoryInterface
}

func NewCellUseCases(r repositories.CellRepositoryInterface) CellUseCasesInterface {
	return CellUseCases{
		CellRepository: r,
	}
}

func (uc CellUseCases) GetCell(id uint) (*models.Cell, error) {
	return uc.CellRepository.Find(id)
}

func (uc CellUseCases) ClickCell(cell *models.Cell, game *models.Game) error {
	cell.CountMinesAround(game)

	return uc.CellRepository.Update(cell)
}
