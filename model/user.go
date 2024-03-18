package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=6"`
	Age       int       `json:"age" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=6" json:"password"`
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
