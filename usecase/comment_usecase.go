package usecase

import (
	"hacktiv-assignment-final/model"
	"hacktiv-assignment-final/repository"
)

type CommentUsecase interface {
	Create(comment model.Comment) (model.CreateCommentResponse, error)
	FindAll() ([]model.Comment, error)
	FindById(id int) (model.Comment, error)
	Update(comment model.Comment) (model.UpdateCommentResponse, error)
	Delete(id int) error
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

func (c *commentUsecase) FindById(id int) (model.Comment, error) {
	return c.repository.FindById(id)
}

func (c *commentUsecase) Update(comment model.Comment) (model.UpdateCommentResponse, error) {
	return c.repository.Update(comment)
}

func (c *commentUsecase) Delete(id int) error {
	return c.repository.Delete(id)
}

func NewCommentUsecase(repository repository.CommentRepository) CommentUsecase {
	return &commentUsecase{repository}
}
