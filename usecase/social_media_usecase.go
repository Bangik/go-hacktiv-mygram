package usecase

import (
	"hacktiv-assignment-final/model"
	"hacktiv-assignment-final/repository"
)

type SocialMediaUsecase interface {
	Create(socialMedia model.SocialMedia) (model.CreateSocialMediaResponse, error)
	FindAll() ([]model.SocialMedia, error)
	FindById(id int) (model.SocialMedia, error)
	Update(socialMedia model.SocialMedia) (model.UpdateSocialMediaResponse, error)
	Delete(id int) error
}

type socialMediaUsecase struct {
	socialMediaRepository repository.SocialMediaRepository
}

func (s *socialMediaUsecase) Create(socialMedia model.SocialMedia) (model.CreateSocialMediaResponse, error) {
	return s.socialMediaRepository.Create(socialMedia)
}

func (s *socialMediaUsecase) FindAll() ([]model.SocialMedia, error) {
	return s.socialMediaRepository.FindAll()
}

func (s *socialMediaUsecase) FindById(id int) (model.SocialMedia, error) {
	return s.socialMediaRepository.FindById(id)
}

func (s *socialMediaUsecase) Update(socialMedia model.SocialMedia) (model.UpdateSocialMediaResponse, error) {
	return s.socialMediaRepository.Update(socialMedia)
}

func (s *socialMediaUsecase) Delete(id int) error {
	return s.socialMediaRepository.Delete(id)
}

func NewSocialMediaUsecase(socialMediaRepository repository.SocialMediaRepository) SocialMediaUsecase {
	return &socialMediaUsecase{socialMediaRepository}
}
