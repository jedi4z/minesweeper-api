package container

import "github.com/jedi4z/minesweeper-api/app/usecases"

type Container struct {
	GameUseCases usecases.GameUseCasesInterface
	CellUseCases usecases.CellUseCasesInterface
}

func NewContainer(
	guc usecases.GameUseCasesInterface,
	cuc usecases.CellUseCasesInterface,
) Container {
	return Container{
		GameUseCases: guc,
		CellUseCases: cuc,
	}
}
