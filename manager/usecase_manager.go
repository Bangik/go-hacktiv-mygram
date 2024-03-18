package manager

import "hacktiv-assignment-final/usecase"

type UseCaseManager interface {
	UserUsecase() usecase.UserUsecase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.repoManager.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
