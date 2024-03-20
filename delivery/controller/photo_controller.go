package controller

import (
	"hacktiv-assignment-final/delivery/middleware"
	"hacktiv-assignment-final/model"
	"hacktiv-assignment-final/usecase"
	"hacktiv-assignment-final/utils/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	Router       *gin.Engine
	photoUsecase usecase.PhotoUsecase
}

func (c *PhotoController) Create(ctx *gin.Context) {
	var photo model.Photo
	var photoResponse model.CreatePhotoRequest
	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := security.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	photo.UserId = userId

	photoResponse, err = c.photoUsecase.Create(photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, photoResponse)
}

func (c *PhotoController) FindAll(ctx *gin.Context) {
	userId, err := security.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	photos, err := c.photoUsecase.FindAll(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var PhotosResponse []model.PhotosResponse
	for _, photo := range photos {
		PhotosResponse = append(PhotosResponse, model.PhotosResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserId:    photo.UserId,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: model.UserPhotosResponse{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}

	ctx.JSON(http.StatusOK, PhotosResponse)
}

func NewPhotoController(router *gin.Engine, photoUsecase usecase.PhotoUsecase) {
	controller := PhotoController{
		Router:       router,
		photoUsecase: photoUsecase,
	}

	roterGroup := router.Group("/photos")
	roterGroup.POST("/", middleware.AuthMiddleware(), controller.Create)
	roterGroup.GET("/", middleware.AuthMiddleware(), controller.FindAll)
}
