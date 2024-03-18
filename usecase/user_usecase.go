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
	Update(user model.UpdateUserResquest) (model.UpdateUserResponse, error)
	Delete(id int) error
	CheckEmailExists(email string) error
	CheckUsernameExists(username string) error
	FindById(id int) (model.User, error)
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

func (u *userUsecase) Update(user model.UpdateUserResquest) (model.UpdateUserResponse, error) {
	return u.repository.Update(user)
}

func (u *userUsecase) Delete(id int) error {
	return u.repository.Delete(id)
}

func (u *userUsecase) CheckEmailExists(email string) error {
	return u.repository.CheckEmailExists(email)
}

func (u *userUsecase) CheckUsernameExists(username string) error {
	return u.repository.CheckUsernameExists(username)
}

func (u *userUsecase) FindById(id int) (model.User, error) {
	return u.repository.FindById(id)
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}
