package manager

import "hacktiv-assignment-final/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	PhotoRepo() repository.PhotoRepository
	CommentRepo() repository.CommentRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Connection())
}

func (r *repoManager) PhotoRepo() repository.PhotoRepository {
	return repository.NewPhotoRepository(r.infra.Connection())
}

func (r *repoManager) CommentRepo() repository.CommentRepository {
	return repository.NewCommentRepository(r.infra.Connection())
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}
