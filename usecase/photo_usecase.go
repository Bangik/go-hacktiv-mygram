package usecase

import (
	"hacktiv-assignment-final/model"
	"hacktiv-assignment-final/repository"
)

type PhotoUsecase interface {
	Create(photo model.Photo) (model.CreatePhotoResponse, error)
	FindAll() ([]model.Photo, error)
	FindById(id int) (model.Photo, error)
	Update(photo model.Photo) (model.UpdatePhotoResponse, error)
	Delete(id int) error
}

type photoUsecase struct {
	repository repository.PhotoRepository
}

func (p *photoUsecase) Create(photo model.Photo) (model.CreatePhotoResponse, error) {
	return p.repository.Create(photo)
}

func (p *photoUsecase) FindAll() ([]model.Photo, error) {
	return p.repository.FindAll()
}

func (p *photoUsecase) FindById(id int) (model.Photo, error) {
	return p.repository.FindById(id)
}

func (p *photoUsecase) Update(photo model.Photo) (model.UpdatePhotoResponse, error) {
	return p.repository.Update(photo)
}

func (p *photoUsecase) Delete(id int) error {
	return p.repository.Delete(id)
}

func NewPhotoUsecase(repository repository.PhotoRepository) PhotoUsecase {
	return &photoUsecase{repository}
}
