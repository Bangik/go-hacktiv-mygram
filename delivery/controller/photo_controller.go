package controller

import (
	"hacktiv-assignment-final/delivery/middleware"
	"hacktiv-assignment-final/model"
	"hacktiv-assignment-final/usecase"
	"hacktiv-assignment-final/utils/security"
	"net/http"
	"strconv"

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

func (c *PhotoController) Update(ctx *gin.Context) {
	var photo model.Photo
	var photoResponse model.UpdatePhotoRequest
	idParam := ctx.Param("photoId")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	findPhoto, err := c.photoUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "photo not found"})
		return
	}

	userId, err := security.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if userId != findPhoto.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this photo"})
		return
	}

	photo.ID = id
	photo.UserId = userId
	err = ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photoResponse, err = c.photoUsecase.Update(photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, photoResponse)
}

func (c *PhotoController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("photoId")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	findPhoto, err := c.photoUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "photo not found"})
		return
	}

	userId, err := security.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if userId != findPhoto.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this photo"})
		return
	}

	err = c.photoUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Your photo has been successfully deleted"})
}

func NewPhotoController(router *gin.Engine, photoUsecase usecase.PhotoUsecase) {
	controller := PhotoController{
		Router:       router,
		photoUsecase: photoUsecase,
	}

	roterGroup := router.Group("/photos")
	roterGroup.POST("/", middleware.AuthMiddleware(), controller.Create)
	roterGroup.GET("/", middleware.AuthMiddleware(), controller.FindAll)
	roterGroup.PUT("/:photoId", middleware.AuthMiddleware(), controller.Update)
	roterGroup.DELETE("/:photoId", middleware.AuthMiddleware(), controller.Delete)
}
