package usecases

import (
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jedi4z/minesweeper-api/app/repositories"
	log "github.com/sirupsen/logrus"
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
	game.InitGame()
	return uc.GameRepository.Insert(game)
}

func (uc GameUseCases) ListGames() ([]models.Game, error) {
	return uc.GameRepository.FindAll()
}

func (uc GameUseCases) GetGame(id uint) (*models.Game, error) {
	return uc.GameRepository.Find(id)
}

func (uc GameUseCases) HoldGame(id uint) (*models.Game, error) {
	game, err := uc.GameRepository.Find(id)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	game.HoldGame()

	if err := uc.GameRepository.Update(game); err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return game, nil
}

func (uc GameUseCases) ResumeGame(id uint) (*models.Game, error) {
	game, err := uc.GameRepository.Find(id)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	game.ResumeGame()

	if err := uc.GameRepository.Update(game); err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return game, nil
}

func (uc GameUseCases) FlagCell(game *models.Game, cellID uint) error {
	if game.Status == models.PlayingState {
		err := ErrGameNotPlayable
		log.WithError(err).Error()
		return err
	}

	if err := game.FlagCell(cellID); err != nil {
		log.WithError(err).Error()
		return err
	}

	return uc.GameRepository.Update(game)
}

func (uc GameUseCases) UncoverCell(game *models.Game, cellID uint) error {
	if game.Status != models.PlayingState {
		err := ErrGameNotPlayable
		log.WithError(err).Error()
		return err
	}

	if err := game.UncoverCell(cellID); err != nil {
		return err
	}

	game.CheckIfWon()

	return uc.GameRepository.Update(game)
}
