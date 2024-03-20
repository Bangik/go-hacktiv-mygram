package model

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=6" json:"password"`
}

type Register struct {
	Username string `binding:"required" json:"username"`
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=6" json:"password"`
	Age      int    `binding:"required" json:"age"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type UpdateUserResquest struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Age       int       `json:"age" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserResponse struct {
	ID       int       `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Age      int       `json:"age"`
	UpdateAt time.Time `json:"updated_at"`
}

type UserPhotosResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserCommentsResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserSocialMediasResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
