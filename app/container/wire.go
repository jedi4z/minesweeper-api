//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/jedi4z/minesweeper-api/app/repositories/mysql_repositories"
	"github.com/jedi4z/minesweeper-api/app/usecases"
	"github.com/jinzhu/gorm"
)

func InitializeContainer(db *gorm.DB) Container {
	wire.Build(
		// repositories
		mysql_repositories.NewGameRepository,
		mysql_repositories.NewCellRepository,
		// use Cases
		usecases.NewGameUseCases,
		usecases.NewCellUseCases,
		// container
		NewContainer,
	)

	return Container{}
}
