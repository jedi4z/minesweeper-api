package usecases

import (
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jedi4z/minesweeper-api/app/repositories"
)

type GameUseCases struct {
	GameRepository repositories.GameRepositoryInterface
}

func NewGameUseCases(r repositories.GameRepositoryInterface) GameUseCasesInterface {
	return GameUseCases{
		GameRepository: r,
	}
}

func (uc GameUseCases) CreateGame(game *models.Game) error {
	game.CreateGrid()
	game.SeedMines()

	return uc.GameRepository.Insert(game)
}

func (uc GameUseCases) ListGames() ([]models.Game, error) {
	return uc.GameRepository.FindAll()
}

func (uc GameUseCases) GetGame(id uint) (*models.Game, error) {
	return uc.GameRepository.Find(id)
}

func (uc GameUseCases) GetGameByCellID(cellID uint) (*models.Game, error) {
	return uc.GameRepository.FindOneByCellID(cellID)
}
