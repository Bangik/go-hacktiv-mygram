package repository

import (
	"hacktiv-assignment-final/model"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo model.Photo) (model.CreatePhotoRequest, error)
	FindAll(idUser int) ([]model.Photo, error)
	FindById(id int) (model.Photo, error)
	Update(photo model.Photo) (model.UpdatePhotoRequest, error)
	Delete(id int) error
}

type photoRepository struct {
	db *gorm.DB
}

func (p *photoRepository) Create(photo model.Photo) (model.CreatePhotoRequest, error) {
	err := p.db.Create(&photo).Error
	createPhotoRequest := model.CreatePhotoRequest{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		CreatedAt: photo.CreatedAt,
	}

	return createPhotoRequest, err
}

func (p *photoRepository) FindAll(idUser int) ([]model.Photo, error) {
	var photos []model.Photo
	err := p.db.Preload("User").Where("user_id = ?", idUser).Find(&photos).Error
	if err != nil {
		return nil, err
	}
	return photos, nil
}

func (p *photoRepository) FindById(id int) (model.Photo, error) {
	var photo model.Photo
	err := p.db.Where("id = ?", id).First(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (p *photoRepository) Update(photo model.Photo) (model.UpdatePhotoRequest, error) {
	err := p.db.Save(&photo).Error
	createPhotoRequest := model.UpdatePhotoRequest{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		UpdatedAt: photo.UpdatedAt,
	}

	return createPhotoRequest, err
}

func (p *photoRepository) Delete(id int) error {
	err := p.db.Where("id = ?", id).Delete(&model.Photo{}).Error
	if err != nil {
		return err
	}

	return nil
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db}
}
