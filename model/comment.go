package model

import "time"

type Comment struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"foreignKey:UserId;references:ID" json:"user"`
	Photo     Photo     `gorm:"foreignKey:PhotoId;references:ID" json:"photo"`
}

type CreateCommentRequest struct {
	PhotoId int    `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required,max=255"`
}

type CreateCommentResponse struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentsResponse struct {
	ID        int                   `json:"id"`
	UserId    int                   `json:"user_id"`
	PhotoId   int                   `json:"photo_id"`
	Message   string                `json:"message"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	User      UserCommentsResponse  `json:"user"`
	Photo     PhotoCommentsResponse `json:"photo"`
}
