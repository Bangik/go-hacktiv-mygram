package usecase

import (
	"hacktiv-assignment-final/model"
	"hacktiv-assignment-final/repository"
)

type CommentUsecase interface {
	Create(comment model.Comment) (model.CreateCommentResponse, error)
	FindAll() ([]model.Comment, error)
}

type commentUsecase struct {
	repository repository.CommentRepository
}

func (c *commentUsecase) Create(comment model.Comment) (model.CreateCommentResponse, error) {
	return c.repository.Create(comment)
}

func (c *commentUsecase) FindAll() ([]model.Comment, error) {
	return c.repository.FindAll()
}

func NewCommentUsecase(repository repository.CommentRepository) CommentUsecase {
	return &commentUsecase{repository}
}
