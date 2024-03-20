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

type SocialMediaController struct {
	Router             *gin.Engine
	socialMediaUsecase usecase.SocialMediaUsecase
}

func (s *SocialMediaController) Create(ctx *gin.Context) {
	var socialMedia model.SocialMedia
	var socialMediaResponse model.CreateSocialMediaResponse
	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := security.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	socialMedia.UserId = userId

	socialMediaResponse, err = s.socialMediaUsecase.Create(socialMedia)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, socialMediaResponse)
}

func (s *SocialMediaController) FindAll(ctx *gin.Context) {
	var socialMediaResponse []model.SocialMediaResponse
	var socialMedias []model.SocialMedia
	var err error

	socialMedias, err = s.socialMediaUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, socialMedia := range socialMedias {
		socialMediaResponse = append(socialMediaResponse, model.SocialMediaResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserId:         socialMedia.UserId,
			CreatedAt:      socialMedia.CreatedAt,
			UpdatedAt:      socialMedia.UpdatedAt,
			User: model.UserSocialMediasResponse{
				ID:       socialMedia.User.ID,
				Username: socialMedia.User.Username,
			},
		})
	}

	if len(socialMediaResponse) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "social media not found"})
		return
	}

	socialMediasResponse := model.SocialMediasResponse{
		SocialMedia: socialMediaResponse,
	}

	ctx.JSON(http.StatusOK, socialMediasResponse)
}

func (s *SocialMediaController) Update(ctx *gin.Context) {
	idParam := ctx.Param("socialMediaId")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	findSocialMedia, err := s.socialMediaUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "social media not found"})
		return
	}

	userId, err := security.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if userId != findSocialMedia.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this social media"})
		return
	}

	var socialMedia model.SocialMedia
	var socialMediaResponse model.UpdateSocialMediaResponse

	socialMedia.ID = id
	socialMedia.UserId = userId

	err = ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	socialMediaResponse, err = s.socialMediaUsecase.Update(socialMedia)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, socialMediaResponse)
}

func (s *SocialMediaController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("socialMediaId")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	findSocialMedia, err := s.socialMediaUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "social media not found"})
		return
	}

	userId, err := security.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if userId != findSocialMedia.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this social media"})
		return
	}

	err = s.socialMediaUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Your social media has been successfully deleted"})
}

func NewSocialMediaController(router *gin.Engine, socialMediaUsecase usecase.SocialMediaUsecase) {
	controller := SocialMediaController{
		Router:             router,
		socialMediaUsecase: socialMediaUsecase,
	}

	routerGroup := router.Group("/socialmedias", middleware.AuthMiddleware())
	routerGroup.POST("/", controller.Create)
	routerGroup.GET("/", controller.FindAll)
	routerGroup.PUT("/:socialMediaId", controller.Update)
	routerGroup.DELETE("/:socialMediaId", controller.Delete)
}
