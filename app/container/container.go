package container

import "github.com/jedi4z/minesweeper-api/app/usecases"

type Container struct {
	GameUseCases usecases.GameUseCasesInterface
}

func NewContainer(guc usecases.GameUseCasesInterface) Container {
	return Container{
		GameUseCases: guc,
	}
}
