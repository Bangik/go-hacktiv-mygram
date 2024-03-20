package model

import "time"

type SocialMedia struct {
	ID             int       `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name" binding:"required"`
	SocialMediaUrl string    `json:"social_media_url" binding:"required"`
	UserId         int       `json:"user_id" gorm:"foreignKey:User"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `json:"user"`
}

type CreateSocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type UpdateSocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaResponse struct {
	ID             int                      `json:"id"`
	Name           string                   `json:"name"`
	SocialMediaUrl string                   `json:"social_media_url"`
	UserId         int                      `json:"user_id"`
	CreatedAt      time.Time                `json:"created_at"`
	UpdatedAt      time.Time                `json:"updated_at"`
	User           UserSocialMediasResponse `json:"user"`
}

type SocialMediasResponse struct {
	SocialMedia []SocialMediaResponse `json:"social_medias"`
}
