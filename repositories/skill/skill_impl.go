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

func (r *repository) FindAll() ([]models.Skill, error) {
	var skills []models.Skill

	err := r.db.Find(&skills).Error

	return skills, err
}

func (r *repository) FindByID(ID int) (models.Skill, error) {
	var skill models.Skill

	err := r.db.Find(&skill, ID).Error

	return skill, err
}

func (r *repository) Create(skill models.Skill) (models.Skill, error) {
	err := r.db.Create(&skill).Error

	return skill, err
}

func (r *repository) Update(skill models.Skill) (models.Skill, error) {
	err := r.db.Save(&skill).Error

	return skill, err
}

func (r *repository) Delete(skill models.Skill) (models.Skill, error) {
	err := r.db.Delete(skill).Error

	return skill, err
}
