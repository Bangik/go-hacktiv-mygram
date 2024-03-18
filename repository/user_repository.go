package repository

import (
	"fmt"
	"hacktiv-assignment-final/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user model.User) (model.RegisterResponse, error)
	Login(user model.User) (model.User, error)
	CheckEmailExists(email string) error
	CheckUsernameExists(username string) error
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Register(user model.User) (model.RegisterResponse, error) {
	var registerResponse model.RegisterResponse
	tx := u.db.Begin()
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return registerResponse, err
	}

	registerResponse = model.RegisterResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Age:      user.Age,
	}

	tx.Commit()
	return registerResponse, nil
}

func (u *userRepository) Login(user model.User) (model.User, error) {
	panic("unimplemented")
}

func (u *userRepository) CheckEmailExists(email string) error {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	fmt.Println(err)
	if err != nil {
		return nil
	}

	return fmt.Errorf("email already exists")
}

func (u *userRepository) CheckUsernameExists(username string) error {
	var user model.User
	err := u.db.Where("username = ?", username).First(&user).Error
	fmt.Println(err)
	if err != nil {
		return nil
	}

	return fmt.Errorf("username already exists")
}

func (u *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
