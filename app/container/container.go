package container

import "github.com/jedi4z/minesweeper-api/app/usecases"

type Container struct {
	UserUseCases usecases.UserUseCaseInterface
	GameUseCases usecases.GameUseCasesInterface
}

func NewContainer(
	uuc usecases.UserUseCaseInterface,
	guc usecases.GameUseCasesInterface,
) Container {
	return Container{
		UserUseCases: uuc,
		GameUseCases: guc,
	}
}
