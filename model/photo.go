package model

import "time"

type Photo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" binding:"required"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"foreignKey:UserId;references:ID"`
}

type CreatePhotoRequest struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdatePhotoRequest struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotosResponse struct {
	ID        int                `json:"id"`
	Title     string             `json:"title"`
	Caption   string             `json:"caption"`
	PhotoUrl  string             `json:"photo_url"`
	UserId    int                `json:"user_id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	User      UserPhotosResponse `json:"user"`
}
