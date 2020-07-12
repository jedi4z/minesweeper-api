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

func (uc GameUseCases) ListGames(user *models.User) ([]models.Game, error) {
	return uc.GameRepository.FindAll(user)
}

func (uc GameUseCases) GetGame(user *models.User, id uint) (*models.Game, error) {
	return uc.GameRepository.Find(user, id)
}

func (uc GameUseCases) HoldGame(game *models.Game) error {
	game.HoldGame()

	if err := uc.GameRepository.Update(game); err != nil {
		log.WithError(err).Error()
		return err
	}

	return nil
}

func (uc GameUseCases) ResumeGame(game *models.Game) error {
	game.ResumeGame()

	if err := uc.GameRepository.Update(game); err != nil {
		log.WithError(err).Error()
		return err
	}

	return nil
}

func (uc GameUseCases) FlagCell(game *models.Game, cellID uint) error {
	if game.Status == models.PlayingState {
		err := errGameNotPlayable
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
		err := errGameNotPlayable
		log.WithError(err).Error()
		return err
	}

	if err := game.UncoverCell(cellID); err != nil {
		return err
	}

	game.CheckIfWon()

	return uc.GameRepository.Update(game)
}
