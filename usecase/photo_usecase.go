package usecase

import (
	"hacktiv-assignment-final/model"
	"hacktiv-assignment-final/repository"
)

type PhotoUsecase interface {
	Create(photo model.Photo) (model.CreatePhotoRequest, error)
	FindAll(idUser int) ([]model.Photo, error)
}

type photoUsecase struct {
	repository repository.PhotoRepository
}

func (p *photoUsecase) Create(photo model.Photo) (model.CreatePhotoRequest, error) {
	return p.repository.Create(photo)
}

func (p *photoUsecase) FindAll(idUser int) ([]model.Photo, error) {
	return p.repository.FindAll(idUser)
}

func NewPhotoUsecase(repository repository.PhotoRepository) PhotoUsecase {
	return &photoUsecase{repository}
}
