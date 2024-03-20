package manager

import "hacktiv-assignment-final/usecase"

type UseCaseManager interface {
	UserUsecase() usecase.UserUsecase
	PhotoUsecase() usecase.PhotoUsecase
	CommnetUsecase() usecase.CommentUsecase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.repoManager.UserRepo())
}

func (u *useCaseManager) PhotoUsecase() usecase.PhotoUsecase {
	return usecase.NewPhotoUsecase(u.repoManager.PhotoRepo())
}

func (u *useCaseManager) CommnetUsecase() usecase.CommentUsecase {
	return usecase.NewCommentUsecase(u.repoManager.CommentRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
