package repository

import (
	"hacktiv-assignment-final/model"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia model.SocialMedia) (model.CreateSocialMediaResponse, error)
	FindAll() ([]model.SocialMedia, error)
	FindById(id int) (model.SocialMedia, error)
	Update(socialMedia model.SocialMedia) (model.UpdateSocialMediaResponse, error)
	Delete(id int) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func (s *socialMediaRepository) Create(socialMedia model.SocialMedia) (model.CreateSocialMediaResponse, error) {
	err := s.db.Create(&socialMedia).Error
	createSocialMediaResponse := model.CreateSocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		CreatedAt:      socialMedia.CreatedAt,
	}

	return createSocialMediaResponse, err
}

func (s *socialMediaRepository) FindAll() ([]model.SocialMedia, error) {
	var socialMedias []model.SocialMedia
	err := s.db.Preload("User").Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}
	return socialMedias, nil
}

func (s *socialMediaRepository) FindById(id int) (model.SocialMedia, error) {
	var socialMedia model.SocialMedia
	err := s.db.Preload("User").First(&socialMedia, id).Error
	if err != nil {
		return model.SocialMedia{}, err
	}
	return socialMedia, nil
}

func (s *socialMediaRepository) Update(socialMedia model.SocialMedia) (model.UpdateSocialMediaResponse, error) {
	err := s.db.Save(&socialMedia).Error
	updateSocialMediaResponse := model.UpdateSocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		UpdatedAt:      socialMedia.UpdatedAt,
	}

	return updateSocialMediaResponse, err
}

func (s *socialMediaRepository) Delete(id int) error {
	err := s.db.Delete(&model.SocialMedia{}, id).Error
	return err
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db}
}
