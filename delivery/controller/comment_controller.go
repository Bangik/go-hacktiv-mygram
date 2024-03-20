package controller

import (
	"hacktiv-assignment-final/delivery/middleware"
	"hacktiv-assignment-final/model"
	"hacktiv-assignment-final/usecase"
	"hacktiv-assignment-final/utils/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	Router         *gin.Engine
	commentUsecase usecase.CommentUsecase
	photoUsecase   usecase.PhotoUsecase
}

func (c *CommentController) Create(ctx *gin.Context) {
	var comment model.Comment
	var commentRequest model.CreateCommentRequest
	var commentResponse model.CreateCommentResponse
	err := ctx.ShouldBindJSON(&commentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := security.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	_, err = c.photoUsecase.FindById(commentRequest.PhotoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	comment.UserId = userId
	comment.PhotoId = commentRequest.PhotoId
	comment.Message = commentRequest.Message

	commentResponse, err = c.commentUsecase.Create(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, commentResponse)
}

func (c *CommentController) FindAll(ctx *gin.Context) {
	comments, err := c.commentUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var commentsResponse []model.CommentsResponse
	for _, comment := range comments {
		commentsResponse = append(commentsResponse, model.CommentsResponse{
			ID:        comment.ID,
			UserId:    comment.UserId,
			PhotoId:   comment.PhotoId,
			Message:   comment.Message,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User: model.UserCommentsResponse{
				ID:       comment.User.ID,
				Username: comment.User.Username,
				Email:    comment.User.Email,
			},
			Photo: model.PhotoCommentsResponse{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserId:   comment.Photo.UserId,
			},
		})
	}

	ctx.JSON(http.StatusOK, commentsResponse)
}

func NewCommentController(router *gin.Engine, commentUsecase usecase.CommentUsecase, photoUsecase usecase.PhotoUsecase) {
	controller := CommentController{
		Router:         router,
		commentUsecase: commentUsecase,
		photoUsecase:   photoUsecase,
	}

	routerGroup := router.Group("/comments")
	routerGroup.POST("/", middleware.AuthMiddleware(), controller.Create)
	routerGroup.GET("/", middleware.AuthMiddleware(), controller.FindAll)
}
