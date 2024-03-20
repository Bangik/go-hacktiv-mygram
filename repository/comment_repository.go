package repository

import (
	"hacktiv-assignment-final/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment model.Comment) (model.CreateCommentResponse, error)
	FindAll() ([]model.Comment, error)
	FindById(id int) (model.Comment, error)
	Update(comment model.Comment) (model.UpdateCommentResponse, error)
	Delete(id int) error
}

type commentRepository struct {
	db *gorm.DB
}

func (c *commentRepository) Create(comment model.Comment) (model.CreateCommentResponse, error) {
	err := c.db.Create(&comment).Error
	createCommentResponse := model.CreateCommentResponse{
		ID:        comment.ID,
		UserId:    comment.UserId,
		PhotoId:   comment.PhotoId,
		Message:   comment.Message,
		CreatedAt: comment.CreatedAt,
	}

	return createCommentResponse, err
}

func (c *commentRepository) FindAll() ([]model.Comment, error) {
	var comments []model.Comment
	err := c.db.Preload("User").Preload("Photo").Find(&comments).Error
	return comments, err
}

func (c *commentRepository) FindById(id int) (model.Comment, error) {
	var comment model.Comment
	err := c.db.First(&comment, id).Error
	return comment, err
}

func (c *commentRepository) Update(comment model.Comment) (model.UpdateCommentResponse, error) {
	err := c.db.Save(&comment).Error
	createCommentResponse := model.UpdateCommentResponse{
		ID:        comment.ID,
		UserId:    comment.UserId,
		PhotoId:   comment.PhotoId,
		Message:   comment.Message,
		UpdatedAt: comment.UpdatedAt,
	}

	return createCommentResponse, err
}

func (c *commentRepository) Delete(id int) error {
	var comment model.Comment
	err := c.db.Delete(&comment, id).Error
	return err
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}
