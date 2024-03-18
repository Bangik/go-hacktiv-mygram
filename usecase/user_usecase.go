package usecase

import (
	"errors"
	"hacktiv-assignment-final/model"
	"hacktiv-assignment-final/repository"
	"hacktiv-assignment-final/utils/security"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(user model.User) (model.RegisterResponse, error)
	Login(user model.Login) (string, error)
	CheckEmailExists(email string) error
	CheckUsernameExists(username string) error
}

type userUsecase struct {
	repository repository.UserRepository
}

func (u *userUsecase) Register(user model.User) (model.RegisterResponse, error) {
	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.RegisterResponse{}, err
	}

	userModel := model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(bytesPassword),
		Age:      user.Age,
	}

	return u.repository.Register(userModel)
}

func (u *userUsecase) Login(user model.Login) (string, error) {
	userModel, err := u.repository.FindByEmail(user.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := security.CreateAccessToken(userModel)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *userUsecase) CheckEmailExists(email string) error {
	return u.repository.CheckEmailExists(email)
}

func (u *userUsecase) CheckUsernameExists(username string) error {
	return u.repository.CheckUsernameExists(username)
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}
