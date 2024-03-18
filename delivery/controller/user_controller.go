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

type UserController struct {
	Router      *gin.Engine
	userUsecase usecase.UserUsecase
}

func (c *UserController) Register(ctx *gin.Context) {
	var register model.User
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.userUsecase.CheckEmailExists(register.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	err = c.userUsecase.CheckUsernameExists(register.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	if register.Age < 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "age must be greater than 8"})
		return
	}

	user, err := c.userUsecase.Register(register)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Login(ctx *gin.Context) {
	var login model.Login
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.userUsecase.Login(login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *UserController) Update(ctx *gin.Context) {
	var user model.UpdateUserResquest
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	_, err = c.userUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	userId, err := security.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if userId != id {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this user"})
		return
	}

	user.ID = id
	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRes, err := c.userUsecase.Update(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userRes)
}

func NewUserController(router *gin.Engine, userUsecase usecase.UserUsecase) *UserController {
	controller := &UserController{
		Router:      router,
		userUsecase: userUsecase,
	}

	roterGroup := router.Group("/users")
	roterGroup.POST("/register", controller.Register)
	roterGroup.POST("/login", controller.Login)
	roterGroup.PUT("/:id", middleware.AuthMiddleware(), controller.Update)

	return controller
}
