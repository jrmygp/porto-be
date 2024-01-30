package repositories

import (
	"porto-be/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Tech, error) {
	var techs []models.Tech

	err := r.db.Find(&techs).Error

	return techs, err
}

func (r *repository) FindByID(ID int) (models.Tech, error) {
	var tech models.Tech

	err := r.db.Find(&tech, ID).Error

	return tech, err
}

func (r *repository) Create(tech models.Tech) (models.Tech, error) {
	err := r.db.Create(&tech).Error

	return tech, err
}

func (r *repository) Update(tech models.Tech) (models.Tech, error) {
	err := r.db.Save(&tech).Error

	return tech, err
}

func (r *repository) Delete(tech models.Tech) (models.Tech, error) {
	err := r.db.Delete(tech).Error

	return tech, err
}
